package infra

import (
	"nfgo.ga/nfgo/nconf"
	"nfgo.ga/nfgo/x/nsecurity"
)

// Config -
type Config struct {
	*nconf.Config
	Security         *nsecurity.SecurityConfig `yaml:"security"`
	VerifyReplay     bool                      `yaml:"verifyReplay"`
	VerifySignature  bool                      `yaml:"verifySignature"`
	VerifyTimeWindow bool                      `yaml:"verifyTimeWindow"`
}

// SetConfig -
func (c *Config) SetConfig(config *nconf.Config) {
	c.Config = config
}
