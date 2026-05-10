package usecases

import (
	"context"
	"errors"
	"fmt"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
	"time"
)


type userUsecase struct {
	userRepo domain.UserRepository
	refreshTokenRepo domain.RefreshTokenRepository
	jwtService domain.JWTService
	tokenService domain.TokenService
}

func NewUserUsecase (
	userRepo domain.UserRepository, 
	refresthTokenRepo domain.RefreshTokenRepository, 
	jwtSvc domain.JWTService,
	tokenSvc domain.TokenService,
	) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		refreshTokenRepo: refresthTokenRepo,
		jwtService: jwtSvc,
		tokenService: tokenSvc,
	}
}

func (uu *userUsecase) Register(ctx context.Context, user domain.User) error {
	// Check if username already exists
	_, err := uu.userRepo.GetByUsername(ctx, user.Username)
	if err == nil {
		return fmt.Errorf("Username already exists")
	}
	// Login: First user is admin
	count, err := uu.userRepo.Count(ctx)
	if err != nil {
		return err
	}

	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	hashed, _ := infrastructure.HashPassword(user.Password)
	user.Password = hashed
	return uu.userRepo.Create(ctx, user)
}

func (uu *userUsecase) Login(ctx context.Context, username, password string) (string, string,error) {
	// Find user
	var user domain.User
	user, usrErr := uu.userRepo.GetByUsername(ctx, username)
	if usrErr != nil {
		return "", "", usrErr
	}

	// Match password
	if match := infrastructure.ComparePassword(password, user.Password); !match {
		return "", "", errors.New("Incorrect cridential")
	}

	// Generate token
	accessToken, _ := uu.jwtService.GenerateAccessToken(user)

	// Generate Refresh token (Opaque token/string)
	rawRefreshToken, _ := uu.tokenService.GenerateRandomString(32)

	// Hash
	hashedToken := uu.tokenService.HashToken(rawRefreshToken)
	
	// Store refresh token in DB
	err := uu.refreshTokenRepo.Create(ctx, domain.RefreshToken{
		UserId: user.ID,
		Token: hashedToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	})

	return accessToken, rawRefreshToken, err
}

func (uu *userUsecase) PromoteUser(ctx context.Context, username string) error {
	return uu.userRepo.UpdateRole(ctx, username, "admin")
}

func (uu *userUsecase) Refresh(ctx context.Context, rawToken string) (string, error) {
	// hash incoming token
	hashedToken := uu.tokenService.HashToken(rawToken)

	// check db
	storedToken, err := uu.refreshTokenRepo.GetByToken(ctx, hashedToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// check expiry
	if time.Now().After(storedToken.ExpiresAt) {
		uu.refreshTokenRepo.DeleteByToken(ctx, hashedToken)
		return "", errors.New("refresh token expired")
	}

	// generate new access token
	user, _ := uu.userRepo.GetByObjectID(ctx, storedToken.UserId)
	newAccessToken, _ := uu.jwtService.GenerateAccessToken(user)

	return newAccessToken, nil
}