package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Task struct {
	Taskid         uuid.UUID          `json:"taskId"`
	CategoryType   string             `json:"categoryType"`
	CreatedDate    pgtype.Timestamptz `json:"createDate"`
	LastUpdateDate pgtype.Timestamptz `json:"lastupdateDate"`
	LastUpdateBy   string             `json:"lastupdateBy"`
	Price          int32              `json:"price"`
}

type TaskUsecase interface {
	CreateTask(Task) (Task, error)
}

type TaskRepository interface {
	CreateTask(Task) (Task, error)
}
