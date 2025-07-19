package utils

import (
	"time"

	"github.com/alifrahmadian/habit-tracker-app-backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RoleId   int64  `json:"role_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, secretKey string, ttl int) (string, error) {
	jwtSecret := []byte(secretKey)

	claims := Claims{
		Id:       user.Id.String(),
		Username: user.Username,
		RoleId:   user.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(ttl))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
