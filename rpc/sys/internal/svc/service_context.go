package svc

import (
	"adminv-api/rpc/model/sys_model"
	"adminv-api/rpc/sys/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel     sys_model.SysUserModel
	LoginLogModel sys_model.SysLoginLogModel
	SysLogModel   sys_model.SysLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Infof("c.Mysql.DataSource: %s", c.Mysql.DataSource)
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     sys_model.NewSysUserModel(sqlConn),
		LoginLogModel: sys_model.NewSysLoginLogModel(sqlConn),
		SysLogModel:   sys_model.NewSysLogModel(sqlConn),
	}

}
