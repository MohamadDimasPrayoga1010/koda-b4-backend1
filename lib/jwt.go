package lib

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type UserPayload struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int) (string, error) {
	_ = godotenv.Load()

	secretKey := os.Getenv("JWT_SECRET")

	claims := UserPayload{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-jwt-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		tokenString, _ := strings.CutPrefix(authHeader, "Bearer ")

		_ = godotenv.Load()
		secret := os.Getenv("JWT_SECRET")

		claims := &UserPayload{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.Id)
		ctx.Next()
	}
}
