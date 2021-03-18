package auth

import (
	"nfgo.ga/nfgo/nutil/nconst"
	"nfgo.ga/nfgo/web"
	"nfgo.ga/nfgo/x/nsecurity"
)

// AuthcAndAuthz -
func (a *AuthAPI) AuthcAndAuthz(c *web.Context) {

	annonAllowed, err := a.enforcer.Enforce("anonymous", c.Request.RequestURI, c.Request.Method)
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
	if err := a.authService.Authenticate(c, ticket); err != nil {
		c.Fail(err)
		c.Abort()
		return
	}

	// 授权
	if err = a.authService.Authorize(c, ticket.Subject, c.Request.RequestURI, c.Request.Method); err != nil {
		c.Fail(err)
		c.Abort()
		return
	}

	c.Next()

}
