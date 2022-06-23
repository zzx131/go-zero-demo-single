package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo-single/app/usercenter/cmd/api/internal/config"
	"go-zero-demo-single/app/usercenter/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.SysUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	db, _ := sqlConn.RawDB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewSysUserModel(sqlConn, c.Cache),
	}
}
