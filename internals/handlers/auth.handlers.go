package handlers

import (
	"bookstoreGo/internals/models"
	"bookstoreGo/internals/repositories"
	"bookstoreGo/pkg"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	*repositories.AuthRepo
}

func InitAuthHandler(a *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{a}
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	//ambil body
	body := &models.AuthModel{}
	if err := ctx.ShouldBind(body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := a.FindByEmail(*body)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	//pengecekan duplicate
	if len(result) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is already registered",
		})
		return
	}

	hash, err := argon2id.CreateHash(body.Password, argon2id.DefaultParams)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := a.SaveUser(models.AuthModel{
		//Id:       body.Id,
		Email:    body.Email,
		Password: hash,
	}); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success registered",
	})
}
func (a *AuthHandler) Login(ctx *gin.Context) {
	body := models.AuthModel{} //kurung kurawal berarti inisialisasi
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := a.FindByEmail(body)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if len(result) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email is not registered",
		})
		return
	}
	match, err := argon2id.ComparePasswordAndHash(body.Password, result[0].Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !match {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bad Credentials",
		})
		return
	}

	payload := pkg.NewPayload(body.Email)
	token, err := payload.CreateToken()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}
