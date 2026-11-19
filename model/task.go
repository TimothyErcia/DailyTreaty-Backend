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
	ErrMissingTaskID = errors.New("task id is missing")
)

// TaskPayload Model
type TaskPayload struct {
	CategoryType string `json:"categoryType"`
	LastUpdateBy string `json:"lastupdateBy"`
	Price        int32  `json:"price"`
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
	CreateTask(context context.Context, taskPayload *TaskPayload) error
	DeleteTaskByID(context context.Context, taskID *uuid.UUID) (*[]Task, error)
	DeleteTaskByCategory(context context.Context, categoryType *string) (*[]Task, error)
}

type TaskUsecase interface {
	CreateTask(context context.Context, taskpayload *TaskPayload) error
	DeleteTaskByID(context context.Context, taskID *string) (*[]Task, error)
	DeleteTaskByCategory(context context.Context, categoryType *string) (*[]Task, error)
}
