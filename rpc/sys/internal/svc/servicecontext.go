package svc

import (
	"adminv-api/rpc/model/sysmodel"
	"adminv-api/rpc/sys/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel sysmodel.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Infof("c.Mysql.DataSource: %s", c.Mysql.DataSource)
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: sysmodel.NewSysUserModel(sqlConn),
	}

}
