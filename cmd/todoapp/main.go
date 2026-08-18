package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/haytoku/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/haytoku/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/haytoku/golang-todoapp/internal/core/transport/http/server"
	user_transport_http "github.com/haytoku/golang-todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to create logger:", err)

		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Logger initialized")

	usersTransportHTTP := user_transport_http.NewUserHttpHandler(nil)
	usersRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.APIVersion1)
	apiVersionRouter.RegisterRoutes(usersRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequsetID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterRouter(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run failed", zap.Error(err))

	}
}
