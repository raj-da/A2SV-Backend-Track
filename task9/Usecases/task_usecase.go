package usecases

import (
	"context"
	domain "task-manager/Domain"
)

type taskUsercase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsercase(repo domain.TaskRepository) domain.TaskUsecase {
	return &taskUsercase{
		taskRepo: repo,
	}
}

func (tu *taskUsercase) Create(ctx context.Context, task domain.Task) error {
	return tu.taskRepo.Create(ctx, task)
}

func (tu *taskUsercase) GetByID(ctx context.Context, id string) (domain.Task, error) {
	return tu.taskRepo.GetByID(ctx, id)
}

func (tu *taskUsercase) GetAll(ctx context.Context) ([]domain.Task, error) {
	return tu.taskRepo.GetAll(ctx)
}

func (tu *taskUsercase) Update(ctx context.Context, id string, task domain.Task) error {
	return tu.taskRepo.Update(ctx, id, task)
}

func (tu *taskUsercase) Delete(ctx context.Context, id string) error {
	return tu.taskRepo.Delete(ctx, id)
}