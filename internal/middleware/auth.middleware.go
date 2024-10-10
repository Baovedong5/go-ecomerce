package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/phuong/go-ecomerce/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorReponse(ctx, response.ErrorInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
