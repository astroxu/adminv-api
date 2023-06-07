package logic

import (
	"adminv-api/rpc/model/sys_model"
	"context"
	"time"

	"adminv-api/rpc/sys/internal/svc"
	"adminv-api/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LoginLogAdd 登录日志
func (l *LoginLogAddLogic) LoginLogAdd(in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(l.ctx, &sys_model.SysLoginLog{
		UserName:   in.UserName,
		Status:     in.Status,
		Ip:         in.Ip,
		CreateTime: time.Now(),
		CreateBy:   in.CreateBy,
	})

	if err != nil {
		return nil, err
	}
	return &sys.LoginLogAddResp{}, nil
}
