package auth

import (
	"context"
	"fmt"
	"nfgo-ddd-showcase/internal/infra"
	"nfgo-ddd-showcase/internal/infra/util"
	"strconv"

	"github.com/casbin/casbin/v2"
	"nfgo.ga/nfgo/ncontext"
	"nfgo.ga/nfgo/ndb"
	"nfgo.ga/nfgo/nerrors"
	"nfgo.ga/nfgo/x/nsecurity"
)

// AuthService -
type AuthService interface {
	Login(ctx context.Context, username string, password string) (*AuthUser, error)

	Register(ctx context.Context, user *AuthUser) error

	Authenticate(ctx context.Context, ticket *nsecurity.AuthTicket) error

	Authorize(ctx context.Context, sub string, obj string, act string) error

	LoadAuthzPolicies(ctx context.Context) error

	FindRoles(ctx context.Context, cond *FindRoleCond) ([]*AuthRole, error)

	UploadAvatar(ctx context.Context, userID int64, avatar []byte) (string, error)
}

func NewAuthService(
	transactional ndb.Transactional,
	userRepo AuthUserRepo, roleRepo AuthRoleRepo, cacheRepo CacheRepo,
	replayChecker nsecurity.ReplayChecker, signKeyStore nsecurity.SignKeyStore, enforcer casbin.IEnforcer,
	config *infra.Config) AuthService {
	return &authServiceImpl{
		transactional: transactional,
		authUserRepo:  userRepo,
		authRoleRepo:  roleRepo,
		cacheRepo:     cacheRepo,
		replayChecker: replayChecker,
		signKeyStore:  signKeyStore,
		enforcer:      enforcer,
		config:        config,
	}
}

type authServiceImpl struct {
	transactional ndb.Transactional
	authUserRepo  AuthUserRepo
	authRoleRepo  AuthRoleRepo
	cacheRepo     CacheRepo
	replayChecker nsecurity.ReplayChecker
	signKeyStore  nsecurity.SignKeyStore
	enforcer      casbin.IEnforcer
	config        *infra.Config
}

func (s *authServiceImpl) Login(ctx context.Context, username string, password string) (*AuthUser, error) {

	user, err := s.authUserRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, infra.ErrUsernameOrPwd
	}

	if err = user.Login(username, password); err != nil {
		return nil, err
	}

	roles, err := s.authRoleRepo.FindRolesByUser(ctx, user)
	if err != nil {
		return nil, err
	}
	rs := make([]string, len(roles))
	for i := range roles {
		rs[i] = roles[i].Code
	}
	s.cacheRepo.SetUserRoles(ctx, int(user.ID), rs)

	mdc, err := ncontext.CurrentMDC(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.signKeyStore.Store(mdc.ClientType(), fmt.Sprint(user.ID), user.SignKey); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authServiceImpl) Register(ctx context.Context, user *AuthUser) error {
	checkUser, err := s.authUserRepo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if checkUser != nil {
		return infra.ErrExistUsername
	}

	if err = user.Register(); err != nil {
		return err
	}

	return s.authUserRepo.InsertUser(ctx, user)

}

func (s *authServiceImpl) Authenticate(ctx context.Context, ticket *nsecurity.AuthTicket) error {
	// Check Time Window
	if s.config.VerifyTimeWindow {
		if err := ticket.VerifyTimeWindow(s.config.Security.TimeWindow); err != nil {
			return err
		}
	}

	// Check Replay
	if s.config.VerifyReplay {
		if err := s.replayChecker.VerifyReplay(ticket.RequestID); err != nil {
			return err
		}
	}

	// Check Signatrue
	if s.config.VerifySignature {
		singnKey, err := s.signKeyStore.Get(ticket.ClientType, ticket.Subject)
		if err != nil {
			return err
		}
		if !ticket.VerifySignature(singnKey) {
			return nerrors.ErrUnauthorized
		}
	}

	// Check Token
	return ticket.VerifyToken(util.JwtSecret)

}

func (s *authServiceImpl) Authorize(ctx context.Context, sub string, obj string, act string) error {
	userID, err := strconv.Atoi(sub)
	if err != nil {
		return err
	}
	roles, err := s.cacheRepo.GetUserRoles(ctx, userID)
	if err != nil {
		return err
	}

	for _, role := range roles {
		allowed, err := s.enforcer.Enforce(role, obj, act)
		if err != nil {
			return err
		}
		if allowed {
			return nil
		}
	}

	return nerrors.ErrForbidden
}

func (s *authServiceImpl) LoadAuthzPolicies(ctx context.Context) error {
	roles, err := s.authRoleRepo.FindAllRoles(ctx, true)
	if err != nil {
		return err
	}

	rules := [][]string{}
	for _, role := range roles {
		sub := role.Code
		for _, res := range role.Resources {
			obj := res.URL
			if obj == "" {
				continue
			}
			act := res.Method
			if act == "" {
				act = "*"
			}
			rules = append(rules, []string{sub, obj, act})
		}
	}
	return nsecurity.InitPolicy(s.enforcer, s.config.Security, rules)
}

func (s *authServiceImpl) UploadAvatar(ctx context.Context, userID int64, avatar []byte) (string, error) {
	mdc, err := ncontext.CurrentMDC(ctx)
	if err != nil {
		return "", err
	}
	if mdc.SubjectID() != strconv.Itoa(int(userID)) {
		return "", infra.ErrNotOneSelf
	}

	// ... upload avatar to object storage, and get the url ....
	avatarURL := "https://pic1.zhimg.com/80/v2-8df0e1ada7af09d3c62f2ba5ec4e4266_720w.jpg"

	err = s.authUserRepo.UpdateAvatarImage(ctx, userID, avatarURL)

	return avatarURL, err
}

func (s *authServiceImpl) FindRoles(ctx context.Context, cond *FindRoleCond) ([]*AuthRole, error) {
	return s.authRoleRepo.FindRoles(ctx, cond)
}
