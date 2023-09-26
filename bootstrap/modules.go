package bootstrap

import (
	"github.com/alirezaKhaki/go-gin/controller"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/alirezaKhaki/go-gin/middleware"
	"github.com/alirezaKhaki/go-gin/repository"
	"github.com/alirezaKhaki/go-gin/router"
	"github.com/alirezaKhaki/go-gin/service"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controller.Module,
	router.Module,
	lib.Module,
	service.Module,
	middleware.Module,
	repository.Module,
)
