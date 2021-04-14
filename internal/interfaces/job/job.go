package job

import (
	"nfgo-ddd-showcase/internal/infra"

	"github.com/google/wire"
	"nfgo.ga/nfgo/njob"
)

var ProviderSet = wire.NewSet(NewJobServer, NewDemoJob)

func NewJobServer(config *infra.Config, demoJob *DemoJob) njob.Server {
	jobs := njob.Jobs{
		"demoJob": demoJob,
	}
	return njob.MustNewServer(config.Config, njob.JobsOption(jobs))
}
