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

func (j *DemoJob) Run(ctx context.Context) error {
	rs, err := j.authService.FindRoles(ctx, &entity.FindRoleCond{})
	nlog.Logger(ctx).Info(rs)
	return err
}
