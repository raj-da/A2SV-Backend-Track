package usecases_test

import (
	"context"
	"errors"
	domain "task-manager/Domain"
	usecases "task-manager/Usecases"
	"task-manager/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 1. Setup the suite
type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *mocks.TaskRepository
	usecase domain.TaskUsecase
}

// 2. Intitialize before each test
func (s *TaskUsecaseSuite) SetupTest() {
	s.mockRepo = new(mocks.TaskRepository)
	s.usecase = usecases.NewTaskUsercase(s.mockRepo)
}

// --- --- --- --- ---
// Test Cases
// --- --- --- --- ---
func (s *TaskUsecaseSuite) TestCreate_Success() {
	ctx := context.TODO()
	task := domain.Task {
		Title: "Test Task",
		Description: "Testing the usecase",
		Status: "Pending",
	}

	s.mockRepo.On("Create", ctx, task).Return(nil)

	err := s.usecase.Create(ctx, task)

	s.NoError(err)
	s.mockRepo.AssertExpectations(s.T())
}

func (s *TaskUsecaseSuite) TestGetByID_Success() {
	ctx := context.TODO()
	id := "507f1f77bcf86cd799439011"
	expectedTask := domain.Task{
		ID: bson.NewObjectID(),
		Title: "Found Task",
	}

	s.mockRepo.On("GetByID", ctx, id).Return(expectedTask, nil)

	result, err := s.usecase.GetByID(ctx, id)

	s.NoError(err)
	s.Equal(expectedTask.Title, result.Title)
}

func (s *TaskUsecaseSuite) TestGetByID_Error() {
	ctx := context.TODO()
	id := "non-existent"

	s.mockRepo.On("GetByID", ctx, id).Return(domain.Task{}, errors.New("task not found"))

	_, err := s.usecase.GetByID(ctx, id)

	s.Error(err)
	s.Equal("task not found", err.Error())
}

func (s *TaskUsecaseSuite) TestGetAll_Success() {
	ctx := context.TODO()
	taskList := []domain.Task{
		{Title: "Task 1"},
		{Title: "Task 2"},
	}

	s.mockRepo.On("GetAll", ctx).Return(taskList, nil)

	results, err := s.usecase.GetAll(ctx)

	s.NoError(err)
	s.Len(results, 2)
	s.Equal("Task 1", results[0].Title)
}

func (s *TaskUsecaseSuite) TestDelete_Success() {
	ctx := context.TODO()
	id := "123"

	s.mockRepo.On("Delete", ctx, id).Return(nil)

	err := s.usecase.Delete(ctx, id)

	s.NoError(err)
}

// 3. The entry point to run the suite
func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}