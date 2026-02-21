package middleware

import (
	"jiu-tracker/config"
	domain_user "jiu-tracker/src/modules/user/domain"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const bearerPrefix = "Bearer "

func RequireAuth(c *gin.Context) {
	tokenString := ""

	// Prefer Authorization header (e.g. from mobile/SPA: "Bearer <token>")
	if h := c.GetHeader("Authorization"); h != "" {
		tokenString = strings.TrimSpace(h)
		if strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = strings.TrimSpace(tokenString[len(bearerPrefix):])
		}
	}

	// Fall back to cookie (e.g. same-origin web)
	if tokenString == "" {
		var err error
		tokenString, err = c.Cookie("Authorization")
		if err != nil || tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if strings.HasPrefix(tokenString, bearerPrefix) {
			tokenString = strings.TrimSpace(tokenString[len(bearerPrefix):])
		}
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		log.Printf("jwt parse: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user with token sub
		var user domain_user.User
		config.DB.Where("id = ?", claims["sub"]).First(&user)

		if user.ID == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
