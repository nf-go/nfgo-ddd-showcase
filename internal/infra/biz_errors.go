package infra

import (
	"nfgo.ga/nfgo/nerrors"
)

var (

	// ErrUsernameOrPwd -
	ErrUsernameOrPwd = nerrors.NewBizError(10001, "用户名或密码错误")
	// ErrExistUsername -
	ErrExistUsername = nerrors.NewBizError(10002, "用户名已经存在")
	// ErrLackFindCond -
	ErrLackFindCond = nerrors.NewBizError(10003, "缺少查询条件")
	// ErrNotOneSelf -
	ErrNotOneSelf = nerrors.NewBizError(10004, "非本人")
)
