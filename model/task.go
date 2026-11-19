package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Error Types
var (
	ErrNotFound      = errors.New("task not found")
	ErrCategoryType  = errors.New("category not found")
	ErrInvalidPrice  = errors.New("task price is invalid")
	ErrUpdateFailed  = errors.New("task failed to update")
	ErrMissingTaskID = errors.New("task id is missing")
)

// TaskPayload Model
type TaskPayload struct {
	CategoryType string `json:"categoryType"`
	LastUpdateBy string `json:"lastupdateBy"`
	Price        int32  `json:"price"`
}

// TaskUpdatePayload Model
type TaskUpdatePayload struct {
	Taskid string `json:"taskId"`
	TaskPayload
}

// Task Model
type Task struct {
	Taskid         uuid.UUID          `json:"taskId"`
	CategoryType   string             `json:"categoryType"`
	CreatedDate    pgtype.Timestamptz `json:"createDate"`
	LastUpdateDate pgtype.Timestamptz `json:"lastupdateDate"`
	LastUpdateBy   string             `json:"lastupdateBy"`
	Price          int32              `json:"price"`
}

type TaskDeletePayload struct {
	Taskid       string `json:"taskId"`
	TaskCategory string `json:"taskCategory"`
	LastUpdateBy string `json:"lastupdateBy"`
}

type TaskRepository interface {
	CreateTask(context context.Context, taskPayload *TaskPayload) (*Task, error)
	UpdateTask(context context.Context, taskID *uuid.UUID, newTask *TaskPayload) error
	DeleteTaskByID(context context.Context, taskID *uuid.UUID) (*[]Task, error)
	DeleteTaskByCategory(context context.Context, categoryType *string) (*[]Task, error)
}

type TaskUsecase interface {
	CreateTask(context context.Context, taskpayload *TaskPayload) (*Task, error)
	UpdateTask(context context.Context, taskID *string, newTask *TaskPayload) error
	DeleteTaskByID(context context.Context, taskID *string) (*[]Task, error)
	DeleteTaskByCategory(context context.Context, categoryType *string) (*[]Task, error)
}
