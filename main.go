package main

import (
	"context"
	"fmt"

	"github.com/alirezaKhaki/go-gin/controller"
	"github.com/alirezaKhaki/go-gin/lib"
	"github.com/alirezaKhaki/go-gin/middleware"
	"github.com/alirezaKhaki/go-gin/repository"
	"github.com/alirezaKhaki/go-gin/router"
	"github.com/alirezaKhaki/go-gin/service"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	_ = godotenv.Load()
	var CommonModules = fx.Options(
		controller.Module,
		router.Module,
		lib.Module,
		service.Module,
		middleware.Module,
		repository.Module,
	)

	bootStrap(CommonModules)
}

func bootStrap(opt fx.Option) {
	logger := lib.GetLogger()
	opts := fx.Options(
		fx.WithLogger(func() fxevent.Logger {
			return logger.GetFxLogger()
		}),
		fx.Invoke(Run()),
	)
	ctx := context.Background()
	app := fx.New(opt, opts)
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		fmt.Println(err)
		logger.Fatal(err)
	}
}

func Run() lib.CommandRunner {
	return func(
		middleware middleware.Middlewares,
		env lib.Env,
		router lib.RequestHandler,
		route router.Routes,
		logger lib.Logger,
		database lib.Database,
		
	) {
		middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}
