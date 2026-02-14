package infrastructure

import (
	"os"
	domain "task-manager/Domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//* --- --- --- --- --- ---//
//*     Claim Model		   //
//* --- --- --- --- --- ---//

type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Rajaf Dereje",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	return token.SignedString(jwtSecret)
}