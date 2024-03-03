package handlers

import (
	"bookstoreGo/internals/models"
	"bookstoreGo/internals/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*repositories.BookRepo
}

func InitBookhandler(b *repositories.BookRepo) *BookHandler {
	return &BookHandler{b}
}
func (b *BookHandler) GetBooks(ctx *gin.Context) {
	result, err := b.FindAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get book",
		"data":    result,
	})
}
func (b *BookHandler) GetBookById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	result, err := b.FindById(id)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Book Not Found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success get book",
		"data":    result,
	})

}

func (b *BookHandler) DeleteBookById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	result, err := b.FindById(id)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}
	if err := b.DeleteById(id); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"messages": "Success Deleted Book",
	})

}

func (b *BookHandler) CreateBook(ctx *gin.Context) {
	body := models.BookModel{}
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := b.SaveBook(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := b.FindAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success save book",
		"data":    result,
	})
}

func (b *BookHandler) UpdateBookById(ctx *gin.Context) {
	// ambil path variabel dengan nama id, dan konversi ke integer
	id, _ := strconv.Atoi(ctx.Param("id"))

	// buat struct body untuk menampung request dari body
	body := models.BookModel{}
	// ambil body,konversi dari json atau form ke struct
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// cek apakah field pada body ada isinya atau tidak
	if body.Title == "" && (body.Description == nil || *body.Description != "") && body.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "empty field, At least one field must be provided",
		})
		return
	}

	// cari buku berdasarkan id
	result, err := b.FindById(id)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// jika buku tidak ditemukan return not found
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messages": "book not found",
		})
		return
	}

	// update buku berdasarkan id
	if err := b.UpdateById(id, body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// kirim response dalam bentuk json, gin.H untuk membuat map dengan key string & vlaue any
	ctx.JSON(http.StatusOK, gin.H{
		"messages": "success update book",
	})
}
