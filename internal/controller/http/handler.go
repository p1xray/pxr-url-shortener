package http

import (
	"github.com/gin-gonic/gin"
	"github.com/p1xray/pxr-url-shortener/internal/controller/http/base"
	v1 "github.com/p1xray/pxr-url-shortener/internal/controller/http/v1"
	"github.com/p1xray/pxr-url-shortener/internal/server"
)

// Handler is handler for http server requests.
type Handler struct {
	service server.URLService
}

// New creates a new http server request handler.
func New(service server.URLService) *Handler {
	return &Handler{
		service: service,
	}
}

// Init initializes the http server request handler.
func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.initBase(router)
	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	v1Handler := v1.New()
	api := router.Group("/api")
	{
		v1Handler.Init(api)
	}
}

func (h *Handler) initBase(router *gin.Engine) {
	baseHandler := base.New(h.service)
	baseHandler.Init(router)
}
