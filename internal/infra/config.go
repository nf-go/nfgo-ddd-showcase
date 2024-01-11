package infra

import (
	"github.com/nf-go/nfgo/nconf"
	"github.com/nf-go/nsecurity"
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
