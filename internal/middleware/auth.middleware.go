package middleware

import (
	"khalifgfrz/coffee-shop-be-go/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthJwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := pkg.NewResponse(ctx)
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			response.Unauthorized("Unauthorized", nil)
			return
		}

		if !strings.Contains(header, "Bearer") {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		// Bearer Bearer token
		token := strings.Replace(header, "Bearer ", "", -1)

		check, err := pkg.VerifyToken(token)
		if err != nil {
			response.Unauthorized("Inavlid Bearer Token", nil)
			return
		}

		ctx.Set("user_id", check.Id)
		ctx.Set("email", check.Email)
		ctx.Set("role", check.Role)
		ctx.Next()
	}
}