package repo

import (
	"context"

	"RNTreaty/db/dailytreaty"
	"RNTreaty/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		db,
	}
}

// CreateTask implements [model.TaskRepository].
func (repo *TaskRepository) CreateTask(context context.Context, taskPayload *model.TaskPayload) error {
	query := dailytreaty.New(repo.db)
	params := dailytreaty.CreateTaskParams{
		CategoryType: taskPayload.CategoryType,
		LastUpdateBy: taskPayload.LastUpdateBy,
		Price:        taskPayload.Price,
	}
	_, taskErr := query.CreateTask(context, params)
	if taskErr != nil {
		return taskErr
	}

	return nil
}
