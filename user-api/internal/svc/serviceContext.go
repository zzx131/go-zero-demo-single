package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo-single/user-api/internal/config"
	"go-zero-demo-single/user-api/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewSysUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
