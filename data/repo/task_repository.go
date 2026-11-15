package repo

import (
	"context"

	"RNTreaty/db/dailytreaty"
	"RNTreaty/model"

	"github.com/google/uuid"
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
func (repo *TaskRepository) CreateTask(context context.Context, taskPayload *model.TaskPayload) (*model.Task, error) {
	query := dailytreaty.New(repo.db)
	params := dailytreaty.CreateTaskParams{
		CategoryType: taskPayload.CategoryType,
		LastUpdateBy: taskPayload.LastUpdateBy,
		Price:        taskPayload.Price,
	}
	createdTask, taskErr := query.CreateTask(context, params)
	if taskErr != nil {
		return nil, taskErr
	}

	parsedTask := model.Task(createdTask)
	return &parsedTask, nil
}

// UpdateTask implements [model.TaskRepository].
func (repo *TaskRepository) UpdateTask(context context.Context, taskID *uuid.UUID, newTask *model.TaskPayload) error {
	query := dailytreaty.New(repo.db)
	params := dailytreaty.UpdateTaskParams{
		Taskid:       *taskID,
		CategoryType: newTask.CategoryType,
		LastUpdateBy: newTask.LastUpdateBy,
		Price:        newTask.Price,
	}

	return query.UpdateTask(context, params)
}
