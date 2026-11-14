package usecase

import (
	"context"

	"RNTreaty/model"
)

type TaskUsecase struct {
	repo model.TaskRepository
}

func NewTaskUsecase(taskRepository model.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: taskRepository}
}

func (taskUsecase *TaskUsecase) CreateTask(context context.Context, taskPayload *model.TaskPayload) error {
	if taskPayload.CategoryType == "" {
		return model.ErrCategoryType
	}

	if taskPayload.Price <= 0 {
		return model.ErrInvalidPrice
	}

	if taskPayload.LastUpdateBy == "" {
		taskPayload.LastUpdateBy = "Ingested-script"
	}

	if taskErr := taskUsecase.repo.CreateTask(context, taskPayload); taskErr != nil {
		return taskErr
	}

	return nil
}
