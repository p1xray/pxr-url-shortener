package app

import (
	grpcapp "github.com/p1xray/pxr-url-shortener/internal/app/grpc"
	httpapp "github.com/p1xray/pxr-url-shortener/internal/app/http"
	"github.com/p1xray/pxr-url-shortener/internal/config"
	"log/slog"
)

// App is an application.
type App struct {
	GRPCServer *grpcapp.App
	HTTPServer *httpapp.App
}

// New creates a new application.
func New(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	// TODO: create storage instance

	// TODO: create service instance

	grpcApp := grpcapp.New(log, cfg.GRPC.Port)
	httpApp := httpapp.New(log, cfg.HTTP.Port)

	return &App{
		GRPCServer: grpcApp,
		HTTPServer: httpApp,
	}
}
