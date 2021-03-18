package infra

import (
	"nfgo.ga/nfgo/nconf"
	"nfgo.ga/nfgo/x/nsecurity"
)

// Config -
type Config struct {
	*nconf.Config
	Security *nsecurity.SecurityConfig `yaml:"security"`
}

// SetConfig -
func (c *Config) SetConfig(config *nconf.Config) {
	c.Config = config
}
