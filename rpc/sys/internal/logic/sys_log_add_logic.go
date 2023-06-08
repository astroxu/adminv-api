package logic

import (
	"adminv-api/rpc/model/sys_model"
	"context"
	"database/sql"

	"adminv-api/rpc/sys/internal/svc"
	"adminv-api/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 系统操作日志
func (l *SysLogAddLogic) SysLogAdd(in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	_, err := l.svcCtx.SysLogModel.Insert(l.ctx, &sys_model.SysLog{
		UserName:  in.UserName,
		Operation: in.Operation,
		Method:    in.Method,
		Params:    in.Params,
		Time:      in.Time,
		Ip:        sql.NullString{String: in.Ip},
		CreateBy:  in.CreateBy,
	})

	if err != nil {
		return nil, err
	}

	return &sys.SysLogAddResp{}, nil
}
