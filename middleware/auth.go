package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	godotenv.Load()
	var JWTtoken = os.Getenv("JWT_TOKEN")
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")

		tokenString, _ := strings.CutPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(JWTtoken), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{
				"message": "Invalid token",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
