package job

import (
	"context"
	"nfgo-ddd-showcase/internal/domain/auth"

	"github.com/nf-go/nfgo/nlog"
)

type DemoJob struct {
	authService auth.AuthService
}

func NewDemoJob(authService auth.AuthService) *DemoJob {
	return &DemoJob{authService: authService}
}

func (j *DemoJob) Run(ctx context.Context) error {
	rs, err := j.authService.FindRoles(ctx, &auth.FindRoleCond{})
	nlog.Logger(ctx).Info(rs)
	return err
}
