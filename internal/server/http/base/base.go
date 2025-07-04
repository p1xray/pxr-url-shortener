package base

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/p1xray/pxr-url-shortener/internal/server"
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
	shortCode, err := server.GetParamFromRoute(c, "short-code")
	if err != nil {
		server.NotFoundResponse(c, "short code not found")
		return
	}

	longURL, err := r.service.LongURL(c.Request.Context(), shortCode)
	if err != nil {
		if errors.Is(err, storage.ErrEntityNotFound) {
			server.NotFoundResponse(c, "long url not found")
			return
		}

		server.ErrorResponse(c, "internal server error")
		return
	}

	c.Redirect(http.StatusFound, longURL)
}
