package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		message)
}

func NotFoundResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		message)
}
