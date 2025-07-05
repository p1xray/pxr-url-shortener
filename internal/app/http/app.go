package httpapp

import (
	"context"
	"errors"
	"fmt"
	"github.com/p1xray/pxr-url-shortener/internal/lib/logger/sl"
	"github.com/p1xray/pxr-url-shortener/internal/server"
	httpserver "github.com/p1xray/pxr-url-shortener/internal/server/http"
	"log/slog"
	"net/http"
	"time"
)

// App is an HTTP server application.
type App struct {
	log        *slog.Logger
	httpServer *http.Server
}

// New creates new instance of HTTP server application.
func New(log *slog.Logger, addr string, service server.URLService) *App {
	handlers := httpserver.New(service)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: handlers.Init(),
	}

	return &App{
		log:        log,
		httpServer: httpServer,
	}
}

// MustRun runs HTTP server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// Run starts the server.
func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.String("addr", a.httpServer.Addr),
	)

	log.Info("running HTTP server")

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// GracefulStop stops the application.
func (a *App) GracefulStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Stop(ctx); err != nil {
		a.log.Error("HTTP app stop error", sl.Err(err))
	}
}

// Stop stops the server.
func (a *App) Stop(ctx context.Context) error {
	const op = "httpapp.Stop"

	log := a.log.With(
		slog.String("op", op),
		slog.String("addr", a.httpServer.Addr),
	)

	log.Info("shutdowning HTTP server")

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("HTTP server is shutdown")

	return nil
}
