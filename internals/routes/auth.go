package routes

import (
	"bookstoreGo/internals/handlers"
	"bookstoreGo/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAuthRouter(router *gin.Engine, db *sqlx.DB) {
	//subrouter
	authRouter := router.Group("/auth")
	AuthRepo := repositories.InitAuthRepo(db)
	AuthHandler := handlers.InitAuthHandler(AuthRepo)
	//buat rute
	//localhost:8000/auth/new
	authRouter.POST("/new", AuthHandler.Register)
	//login
	authRouter.POST("", AuthHandler.Login)
}
