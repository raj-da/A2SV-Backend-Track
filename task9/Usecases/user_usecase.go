package usecases

import (
	"context"
	"errors"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
)


type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase (repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (uu *userUsecase) Register(ctx context.Context, user domain.User) error {
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

func (uu *userUsecase) Login(ctx context.Context, username, password string) (string, error) {
	// Find user
	var user domain.User
	user, err := uu.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Match password
	if match := infrastructure.ComparePassword(password, user.Password); !match {
		return "", errors.New("Incorrect cridential")
	}

	// Generate token
	token, err := infrastructure.GenerateToken(&user)
	return token, err
}

func (uu *userUsecase) PromoteUser(ctx context.Context, username string) error {
	return uu.userRepo.UpdateRole(ctx, username, "admin")
}