package server

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyParam = errors.New("empty param")
)

func GetParamFromRoute(c *gin.Context, name string) (string, error) {
	routeParam := c.Param(name)
	if routeParam == "" {
		return "", ErrEmptyParam
	}

	return routeParam, nil
}
