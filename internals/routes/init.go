package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	return router
}
