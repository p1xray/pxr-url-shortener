package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/pxr-url-shortener/internal/server/http/v1/ping"
)

// Handler is request handler for API v1.
type Handler struct {
}

// New creates new instance of the API v1 request handler.
func New() *Handler {
	return &Handler{}
}

// Init initializes the API v1 request handler.
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		ping.InitRoutes(v1)
	}
}
