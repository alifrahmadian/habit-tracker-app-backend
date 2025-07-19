package middlewares

import (
	"net/http"
	"strings"
	"time"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/handlers/responses"
	e "github.com/alifrahmadian/habit-tracker-app-backend/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RoleId   int64  `json:"role_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(secretKey string, allowedRoles ...int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.FailedResponse(c, http.StatusUnauthorized, e.ErrNoAuthorizationHeader.Error())
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			responses.FailedResponse(c, http.StatusUnauthorized, e.ErrInvalidTokenFormat.Error())
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return nil, e.ErrUnexpectedSigningMethod
			}

			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			responses.FailedResponse(c, http.StatusUnauthorized, e.ErrTokenExpired.Error())
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			responses.FailedResponse(c, http.StatusUnauthorized, e.ErrTokenExpired.Error())
			c.Abort()
			return
		}

		roleAllowed := false
		for _, role := range allowedRoles {
			if claims.RoleId == role {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			responses.FailedResponse(c, http.StatusUnauthorized, e.ErrAccessDenied.Error())
			c.Abort()
			return
		}

		c.Set("user_id", claims.ID)
		c.Set("username", claims.Username)
		c.Set("role_id", claims.RoleId)

		c.Next()
	}
}
