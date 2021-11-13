package auth

import "github.com/google/wire"

var ProviderSet = wire.NewSet(serviceProviderSet, repoProviderSet)

var serviceProviderSet = wire.NewSet(NewAuthService)

var repoProviderSet = wire.NewSet(NewAuthRoleRepo, NewAuthUserRepo, NewCacheRepo)
