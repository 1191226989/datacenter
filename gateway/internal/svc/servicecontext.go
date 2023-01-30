package svc

import (
	"datacenter/gateway/internal/config"
	"datacenter/gateway/internal/middleware"
	"datacenter/user/rpc/userclient"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CheckPassMiddleware rest.Middleware

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		CheckPassMiddleware: middleware.NewCheckPassMiddleware().Handle,
		UserRpc:             userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
