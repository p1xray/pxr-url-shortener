package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		message)
}

func NotFound(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		message)
}
