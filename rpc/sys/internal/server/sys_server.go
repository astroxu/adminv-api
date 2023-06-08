// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"adminv-api/rpc/sys/internal/logic"
	"adminv-api/rpc/sys/internal/svc"
	"adminv-api/rpc/sys/sys"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

// 登录
func (s *SysServer) Login(ctx context.Context, in *sys.LoginReq) (*sys.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 登录日志
func (s *SysServer) LoginLogAdd(ctx context.Context, in *sys.LoginLogAddReq) (*sys.LoginLogAddResp, error) {
	l := logic.NewLoginLogAddLogic(ctx, s.svcCtx)
	return l.LoginLogAdd(in)
}

func (s *SysServer) LoginLogList(ctx context.Context, in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	l := logic.NewLoginLogListLogic(ctx, s.svcCtx)
	return l.LoginLogList(in)
}

func (s *SysServer) LoginLogDelete(ctx context.Context, in *sys.LoginLogDeleteReq) (*sys.LoginLogDeleteResp, error) {
	l := logic.NewLoginLogDeleteLogic(ctx, s.svcCtx)
	return l.LoginLogDelete(in)
}

// 系统操作日志
func (s *SysServer) SysLogAdd(ctx context.Context, in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	l := logic.NewSysLogAddLogic(ctx, s.svcCtx)
	return l.SysLogAdd(in)
}

func (s *SysServer) SysLogList(ctx context.Context, in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	l := logic.NewSysLogListLogic(ctx, s.svcCtx)
	return l.SysLogList(in)
}

func (s *SysServer) SysLogDelete(ctx context.Context, in *sys.SysLogDeleteReq) (*sys.SysLogDeleteResp, error) {
	l := logic.NewSysLogDeleteLogic(ctx, s.svcCtx)
	return l.SysLogDelete(in)
}
