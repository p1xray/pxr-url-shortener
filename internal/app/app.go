package app

import (
	grpcapp "github.com/p1xray/pxr-url-shortener/internal/app/grpc"
	httpapp "github.com/p1xray/pxr-url-shortener/internal/app/http"
	"github.com/p1xray/pxr-url-shortener/internal/config"
	"github.com/p1xray/pxr-url-shortener/internal/service"
	"github.com/p1xray/pxr-url-shortener/internal/storage/sqlite"
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
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		panic(err)
	}

	urlService := service.New(cfg.ShortCodeGenerator, storage)

	grpcApp := grpcapp.New(log, cfg, urlService)
	httpApp := httpapp.New(log, cfg.HTTP.Addr, urlService)

	return &App{
		GRPCServer: grpcApp,
		HTTPServer: httpApp,
	}
}
