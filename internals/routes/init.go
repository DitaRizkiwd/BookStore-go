package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	InitAuthRouter(router, db)

	return router
}
