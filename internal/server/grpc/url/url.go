package url

import (
	"context"
	"github.com/p1xray/pxr-url-shortener/internal/config"
	"github.com/p1xray/pxr-url-shortener/internal/server"
	"github.com/p1xray/pxr-url-shortener/internal/server/grpc/response"
	urlshortenerpb "github.com/p1xray/pxr-url-shortener/pkg/grpc/gen/go/urlshortener"
	"google.golang.org/grpc"
)

type serverAPI struct {
	urlshortenerpb.UnimplementedUrlShortenerServer
	urlService     server.URLService
	httpServerAddr string
}

// Register registers the implementation of the API service with the gRPC server.
func Register(gRPC *grpc.Server, urlService server.URLService, cfg config.HTTPConfig) {
	urlshortenerpb.RegisterUrlShortenerServer(gRPC, &serverAPI{urlService: urlService, httpServerAddr: cfg.Addr})
}

func (s *serverAPI) Shorten(
	ctx context.Context,
	req *urlshortenerpb.ShortenRequest,
) (*urlshortenerpb.ShortenResponse, error) {
	if err := validateShortenRequest(req); err != nil {
		return nil, err
	}

	shortURL, err := s.urlService.Shorten(ctx, req.GetLongUrl(), s.httpServerAddr)
	if err != nil {
		return nil, response.InternalError("error shortening url")
	}

	return &urlshortenerpb.ShortenResponse{
		ShortCode: shortURL.ShortCode,
		ShortUrl:  shortURL.ShortURL,
	}, nil
}

func validateShortenRequest(req *urlshortenerpb.ShortenRequest) error {
	if req.GetLongUrl() == "" {
		return response.InvalidArgumentError("long url is empty")
	}

	return nil
}
