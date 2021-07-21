// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	template "go-zero-demo/book/service/user/api/internal/handler/template"
	"go-zero-demo/book/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users/login",
				Handler: template.LoginHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}