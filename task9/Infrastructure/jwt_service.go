package infrastructure

import (
	"fmt"
	"os"
	"time"
	domain "task-manager/Domain"
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

type jwtService struct {
	secretKey string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: os.Getenv("JWT_SECRET_KEY"),
	}
} 

func (j *jwtService) GenerateAccessToken(user domain.User) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := Claims {
		Username: user.Username,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "Rajaf Dereje",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateAccessToken(tokenStr string) (*domain.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*domain.TokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}