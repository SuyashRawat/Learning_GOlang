package middleware

import (
	// "fmt"
	"net/http"
	"strings"
	"taskmanage/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		authheader:=ctx.GetHeader("Authorization")

		if authheader==""{
			ctx.JSON(http.StatusUnauthorized,utils.ErrorResponse("missing Authentication Error"))
			ctx.Abort()
			return 
		}
		token:=strings.TrimPrefix(authheader,"Bearer ")
		if token==""{
			ctx.JSON(http.StatusUnauthorized,utils.ErrorResponse("invalid authorization header"))
			ctx.Abort()
			return
		}
		UserId,err:=utils.ValidateToken(token)
		if err!=nil{
			ctx.JSON(http.StatusUnauthorized,utils.ErrorResponse("invalid token"))
			ctx.Abort()
			return
		}

		ctx.Set("UserId",UserId)
		ctx.Next()


	}
}