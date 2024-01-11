package util

import "github.com/nf-go/nfgo/nutil/ntypes"

const (
	// RedisKeyUserRoles - auth:userRoles:{userID}
	RedisKeyUserRoles ntypes.Key = "auth:userRoles:%d"
)
