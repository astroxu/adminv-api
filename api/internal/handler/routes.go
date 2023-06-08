// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	syslog "adminv-api/api/internal/handler/sys/log"
	sysuser "adminv-api/api/internal/handler/sys/user"
	"adminv-api/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/login",
				Handler: sysuser.UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		//rest.WithMiddlewares(
		//[]rest.Middleware{serverCtx.CheckUrl},
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: syslog.SysLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: syslog.SysLogDeleteHandler(serverCtx),
			},
		},
		//}...,
		//),
		//rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/sysLogs"),
	)

	server.AddRoutes(
		//rest.WithMiddlewares(
		//[]rest.Middleware{serverCtx.CheckUrl},
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: syslog.LoginLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: syslog.LoginLogDeleteHandler(serverCtx),
			},
		},
		//}...,
		//),
		//rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/loginLogs"),
	)
}
