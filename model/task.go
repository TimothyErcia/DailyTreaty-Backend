package model

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Error Types
var (
	ErrNotFound     = errors.New("task not found")
	ErrCategoryType = errors.New("category not found")
	ErrInvalidPrice = errors.New("task price is invalid")
)

// TaskPayload Model
type TaskPayload struct {
	CategoryType string
	LastUpdateBy string
	Price        int32
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

type TaskRepository interface {
	CreateTask(context context.Context, taskPayload *TaskPayload) error
}

type TaskUsecase interface {
	CreateTask(context context.Context, taskpayload *TaskPayload) error
}
