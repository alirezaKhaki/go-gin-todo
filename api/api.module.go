package api

import (
	"github.com/alirezaKhaki/go-gin/api/controller"
	"github.com/alirezaKhaki/go-gin/api/middleware"
	"github.com/alirezaKhaki/go-gin/api/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	controller.Module,
	router.Module,
	middleware.Module,
)
