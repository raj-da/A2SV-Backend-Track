package usecases_test

import (
	"context"
	"task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
	usecases "task-manager/Usecases"
	"task-manager/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// 1. Setup the suite
type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.UserRepository
	usecase Domain.UserUsecase
}

// 2. Intitialize before each test
func (s *UserUsecaseSuite) SetupTest() {
	s.mockRepo = new(mocks.UserRepository)
	s.usecase = usecases.NewUserUsecase(s.mockRepo)
}

// --- --- --- --- ---
// Test Cases
// --- --- --- --- ---
func (s *UserUsecaseSuite) TestRegister_Success() {
	ctx := context.TODO()
	user := Domain.User{
		Username: "rajaf",
		Password: "plain123",
		Role: "user",
	}

	var userCount int64 = 1
	s.mockRepo.On("Count", ctx).Return(userCount, nil)
	s.mockRepo.
        On("Create", ctx, mock.MatchedBy(func(u Domain.User) bool {
            return u.Username == "rajaf" &&
                   u.Role == "user" && // because count != 0
                   u.Password != "plain123" // password must be hashed
        })).
        Return(nil)

	err := s.usecase.Register(ctx, user)
	s.NoError(err)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *UserUsecaseSuite) TestLogin_Success() {
	ctx := context.TODO()
	
	hashedPassword, _ := infrastructure.HashPassword("plain123")
	user := Domain.User{
		Username: "rajaf",
		Password: hashedPassword,
	}

	s.mockRepo.On("GetByUsername", ctx, user.Username).Return(user, nil)

	_, err := s.usecase.Login(ctx, user.Username, "plain123")
	s.NoError(err)
}

func (s *UserUsecaseSuite) TestLogin_Error() {
	ctx := context.TODO()
	
	hashedPassword, _ := infrastructure.HashPassword("plain123")
	user := Domain.User{
		Username: "rajaf",
		Password: hashedPassword,
	}

	s.mockRepo.On("GetByUsername", ctx, user.Username).Return(user, nil)

	_, err := s.usecase.Login(ctx, user.Username, "plain1234")
	s.Error(err)
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}