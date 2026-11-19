package usecase

import (
	"context"

	"RNTreaty/model"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo model.TaskRepository
}

func NewTaskUsecase(taskRepository model.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: taskRepository}
}

func (taskUsecase *TaskUsecase) CreateTask(context context.Context, taskPayload *model.TaskPayload) (*model.Task, error) {
	payloadErr := validateTaskPayload(taskPayload)
	if payloadErr != nil {
		return nil, payloadErr
	}

	createdTask, taskErr := taskUsecase.repo.CreateTask(context, taskPayload)
	if taskErr != nil {
		return nil, taskErr
	}

	return createdTask, nil
}

// UpdateTask implements [model.TaskUsecase].
func (taskUsecase *TaskUsecase) UpdateTask(context context.Context, taskID *string, newTask *model.TaskPayload) error {
	if taskID == nil {
		return model.ErrTaskIDMissing
	}

	payloadErr := validateTaskPayload(newTask)
	if payloadErr != nil {
		return payloadErr
	}

	convertedID, parseErr := uuid.Parse(*taskID)
	if parseErr != nil {
		return model.ErrTaskIDMissing
	}

	if taskUpdateErr := taskUsecase.repo.UpdateTask(context, &convertedID, newTask); taskUpdateErr != nil {
		return taskUpdateErr
	}

	return nil
}

func validateTaskPayload(taskPayload *model.TaskPayload) error {
	if taskPayload.CategoryType == "" {
		return model.ErrCategoryType
	}

	if taskPayload.Price <= 0 {
		return model.ErrInvalidPrice
	}

	if taskPayload.LastUpdateBy == "" {
		taskPayload.LastUpdateBy = "Ingested-script"
	}

	return nil
}

// DeleteTaskByID implements [model.TaskUsecase].
func (taskUsecase *TaskUsecase) DeleteTaskByID(context context.Context, taskID *string) (*[]model.Task, error) {
	if taskID == nil {
		return nil, model.ErrMissingTaskID
	}

	convertedID, parseErr := uuid.Parse(*taskID)
	if parseErr != nil {
		return nil, parseErr
	}

	return taskUsecase.repo.DeleteTaskByID(context, &convertedID)
}

// DeleteTaskByCategory implements [model.TaskUsecase].
func (taskUsecase *TaskUsecase) DeleteTaskByCategory(context context.Context, categoryType *string) (*[]model.Task, error) {
	if categoryType == nil {
		return nil, model.ErrMissingTaskID
	}

	return taskUsecase.repo.DeleteTaskByCategory(context, categoryType)
}
