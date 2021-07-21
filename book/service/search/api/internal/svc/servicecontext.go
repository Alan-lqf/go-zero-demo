package svc

import (
	"go-zero-demo/book/service/search/api/internal/config"
	"go-zero-demo/book/service/search/api/internal/middleware"

	"github.com/tal-tech/go-zero/rest"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
