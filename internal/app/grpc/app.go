package grpcapp

import (
	"fmt"
	"github.com/p1xray/pxr-url-shortener/internal/server"
	"log/slog"
	"net"

	urlservice "github.com/p1xray/pxr-url-shortener/internal/server/grpc/url"

	"google.golang.org/grpc"
)

// App is an gRPC application.
type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server application.
func New(
	log *slog.Logger,
	port int,
	service server.URLService,
) *App {
	gRPCServer := grpc.NewServer()
	
	urlservice.Register(gRPCServer, service)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf((":%d"), a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op))
	a.log.Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
