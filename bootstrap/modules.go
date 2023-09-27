package bootstrap

import (
	"github.com/alirezaKhaki/go-gin/api"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/alirezaKhaki/go-gin/repository"
	"github.com/alirezaKhaki/go-gin/service"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	lib.Module,
	service.Module,
	repository.Module,
)
