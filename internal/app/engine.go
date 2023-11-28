package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine interface {
	Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup
	NoRoute(handlers ...gin.HandlerFunc)
	ServeHTTP(http.ResponseWriter, *http.Request)
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
}

func NewEngine() Engine {
	return gin.Default()
}
