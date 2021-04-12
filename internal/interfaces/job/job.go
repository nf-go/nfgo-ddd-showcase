package job

import (
	"github.com/google/wire"
	"nfgo.ga/nfgo/njob"
)

var ProviderSet = wire.NewSet(NewJobs, NewDemoJob)

// Jobs -
func NewJobs(demoJob *DemoJob) njob.Jobs {
	return njob.Jobs{
		"demoJob": demoJob,
	}
}
