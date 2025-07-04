package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/p1xray/pxr-url-shortener/internal/server"
	"github.com/p1xray/pxr-url-shortener/internal/server/http/request"
	http2 "github.com/p1xray/pxr-url-shortener/internal/server/http/response"
	"github.com/p1xray/pxr-url-shortener/internal/storage"
	"net/http"
)

// Routes provides routes for base URL.
type Routes struct {
	service server.URLService
}

// New creates a new instance of routes for base URL.
func New(service server.URLService) *Routes {
	return &Routes{service}
}

// Init initializes routes for base URL.
func (r *Routes) Init(g *gin.Engine) {
	g.GET("/:short-code", r.redirectByShortCode)
}

func (r *Routes) redirectByShortCode(c *gin.Context) {
	shortCode, err := request.GetParamFromRoute(c, "short-code")
	if err != nil {
		http2.NotFound(c, "short code not found")
		return
	}

	longURL, err := r.service.LongURL(c.Request.Context(), shortCode)
	if err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			http2.NotFound(c, "long url not found")
			return
		}

		http2.InternalServerError(c, "internal server error")
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
