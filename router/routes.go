package router

import (
	"net/http"
	"gin-temp-app/constant"
	"gin-temp-app/controller"
)

// this routes represent user
var user = Routes{
	Route{"Hello User", http.MethodPost, constant.LoginUserPath, controller.LoginUser},
	Route{"Hello User", http.MethodGet, constant.HelloUserPath, controller.HelloUser},
}
