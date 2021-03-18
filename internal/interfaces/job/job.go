package job

import (
	"github.com/google/wire"
	"nfgo.ga/nfgo/njob"
)

var ProviderSet = wire.NewSet(NewJobs, NewDemoJob)

// Jobs -
func NewJobs(demoJob *DemoJob) njob.JobFuncs {
	return njob.JobFuncs{
		"demoJob": demoJob.doDemo,
	}
}
