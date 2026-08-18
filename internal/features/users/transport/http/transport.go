package users_transport_http

import (
	"net/http"

	core_http_server "github.com/haytoku/golang-todoapp/internal/core/transport/http/server"
)

type UserHttpHandler struct {
	userService UserService
}

type UserService interface {
}

func NewUserHttpHandler(userService UserService) *UserHttpHandler {
	return &UserHttpHandler{
		userService: userService,
	}
}

func (h *UserHttpHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: http.HandlerFunc(h.CreateUser),
		},
	}

}
