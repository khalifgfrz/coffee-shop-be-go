package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MockAuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := gin.H{}
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			response["message"] = "Unauthorized"
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}

		if header != "Bearer validToken" {
			response["message"] = "Invalid Bearer Token"
			ctx.JSON(http.StatusUnauthorized, response)
			return
		}

		// Simulated token details
		tokenDetails := struct {
			Id    string
			Email string
			Role  string
		}{
			Id:    "123",
			Email: "test@example.com",
			Role:  "admin",
		}

		roleAllowed := false
		for _, role := range roles {
			if tokenDetails.Role == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			response["message"] = "Forbidden: Insufficient permissions"
			ctx.JSON(http.StatusUnauthorized, response)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", tokenDetails.Id)
		ctx.Set("email", tokenDetails.Email)
		ctx.Set("role", tokenDetails.Role)
		ctx.Next()
	}
}