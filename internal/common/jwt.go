package common

import (
	"fmt"
	"go-app/config"
	"go-app/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/net/context"
)

type JwtCustomClaims struct {
	ID       int64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role  	 string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(c context.Context, user *entity.User) (string, error) {
	expiredTime := time.Now().Local().Add(60 * time.Minute)
	claims := JwtCustomClaims{
		ID: user.ID,
		Name:  user.Name,
		Username: user.Username,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var cfg *config.Config
	cfg, _ = config.NewConfig(".env")
	encodedToken, err := token.SignedString([]byte(cfg.JWTSecretKey))

	if err != nil {
		fmt.Println("salah di generate access token fungsi jwt")
		return "", err
	}

	return encodedToken, nil
}