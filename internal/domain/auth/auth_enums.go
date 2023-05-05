//go:generate stringer -type=UserStatus,ResourceType -output enums.string.go
package auth

// UserStatus -
type UserStatus int8

const (
	// UserStatusEnabled -
	UserStatusEnabled UserStatus = 1
	// UserStatusDiabled -
	UserStatusDiabled UserStatus = -1
	// UserStatusLocked -
	UserStatusLocked UserStatus = -2
)

// ResourceType -
type ResourceType int8

const (
	// ResTypeMenu -
	ResTypeMenu ResourceType = 1
)
