package svc

import (
	"github.com/zhoushuguang/lebron/cmd/user/rpc/internal/config"
	"github.com/zhoushuguang/lebron/cmd/user/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn, c.CacheRedis),
	}
}
