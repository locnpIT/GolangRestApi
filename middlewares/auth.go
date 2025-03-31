package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/utils"
)

func Authenticate(context *gin.Context) {

	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Authorization header missing."})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Invalid Authorization header format."})
		return
	}

	tokenString := parts[1]

	userId, err := utils.VerifyToken(tokenString)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Invalid token."})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
