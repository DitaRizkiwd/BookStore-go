package middleware

import (
	"bookstoreGo/pkg"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckToken(ctx *gin.Context) {
	//ambil header authorization
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Please login first",
		})
		return
	}
	if !strings.Contains(bearerToken, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Invalid Authorization",
		})
		return
	}
	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	_, err := pkg.VerifyToken(token)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			log.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Please login again",
			})
			return
		}
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Please login first",
		})
		return
	}
	ctx.Next()
}
