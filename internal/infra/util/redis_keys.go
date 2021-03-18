package util

import "nfgo.ga/nfgo/nutil/ntypes"

const (
	// RedisKeyUserRoles - auth:userRoles:{userID}
	RedisKeyUserRoles ntypes.Key = "auth:userRoles:%d"
)
