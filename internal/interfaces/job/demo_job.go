package job

import (
	"context"
	"nfgo-ddd-showcase/internal/domain/auth/entity"
	"nfgo-ddd-showcase/internal/domain/auth/service"

	"nfgo.ga/nfgo/nlog"
)

type DemoJob struct {
	authService service.AuthService
}

func NewDemoJob(authService service.AuthService) *DemoJob {
	return &DemoJob{authService: authService}
}

func (j *DemoJob) doDemo() {
	rs, err := j.authService.FindRoles(context.Background(), &entity.FindRoleCond{})
	nlog.Info(rs, err)
}
