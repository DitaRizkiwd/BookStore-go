package routes

import (
	"bookstoreGo/internals/handlers"
	"bookstoreGo/internals/middleware"
	"bookstoreGo/internals/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBookRouter(router *gin.Engine, db *sqlx.DB) {
	bookRouter := router.Group("/book")
	bookRepo := repositories.InitBookRepo(db)
	bookHandler := handlers.InitBookhandler(bookRepo)

	//get books
	bookRouter.GET("", bookHandler.GetBooks)

	//create book
	bookRouter.POST("/new", middleware.CheckToken, bookHandler.CreateBook)

	// get book by id
	//localhost:8000/book/id
	bookRouter.GET("/:id", middleware.CheckToken, bookHandler.GetBookById)

	// delete book by id
	//
	bookRouter.DELETE("/:id", middleware.CheckToken, bookHandler.DeleteBookById)

	// update book by id
	bookRouter.PATCH("/:id", middleware.CheckToken, bookHandler.UpdateBookById)
}
