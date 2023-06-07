package log

import (
	"context"

	"adminv-api/api/internal/svc"
	"adminv-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogDeleteLogic) SysLogDelete(req *types.DeleteSysLogReq) (resp *types.DeleteSysLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
