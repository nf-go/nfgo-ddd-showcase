package auth

import (
	"nfgo-ddd-showcase/internal/domain/auth"

	"github.com/casbin/casbin/v2"
	"nfgo.ga/nfgo/nutil/nconst"
	"nfgo.ga/nfgo/web"
	"nfgo.ga/nfgo/x/nsecurity"
)

type MiddleWare struct {
	authService auth.AuthService
	enforcer    casbin.IEnforcer
}

func NewMiddleWare(authService auth.AuthService, enforcer casbin.IEnforcer) *MiddleWare {
	return &MiddleWare{
		authService: authService,
		enforcer:    enforcer,
	}
}

// AuthcAndAuthz -
func (m *MiddleWare) AuthcAndAuthz(c *web.Context) {

	annonAllowed, err := m.enforcer.Enforce("anonymous", c.Request.RequestURI, c.Request.Method)
	if err != nil {
		c.Fail(err)
		c.Abort()
		return
	}
	// 允许匿名访问直接放行
	if annonAllowed {
		c.Next()
		return
	}

	ticket := &nsecurity.AuthTicket{
		ClientType: c.GetHeader(nconst.HeaderClientType),
		RequestID:  c.GetHeader(nconst.HeaderTraceID),
		Token:      c.GetHeader(nconst.HeaderToken),
		Subject:    c.GetHeader(nconst.HeaderSub),
		Timestamp:  c.GetHeader(nconst.HeaderTs),
		Signature:  c.GetHeader(nconst.HeaderSig),
	}
	// 认证
	if err := m.authService.Authenticate(c, ticket); err != nil {
		c.Fail(err)
		c.Abort()
		return
	}

	// 授权
	if err = m.authService.Authorize(c, ticket.Subject, c.Request.URL.Path, c.Request.Method); err != nil {
		c.Fail(err)
		c.Abort()
		return
	}

	c.Next()

}
