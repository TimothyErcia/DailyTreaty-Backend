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

func parseDBTypeToModel(ctx context.Context, qtx *dailytreaty.Queries) (*[]model.Task, error) {
	allTask, allTaskErr := qtx.GetAllTask(ctx)
	if allTaskErr != nil {
		return nil, allTaskErr
	}

	var convertTask []model.Task
	for _, task := range allTask {
		convertTask = append(convertTask, model.Task(task))
	}

	return &convertTask, nil
}

// DeleteTaskByID implements [model.TaskRepository].
func (repo *TaskRepository) DeleteTaskByID(ctx context.Context, taskID *uuid.UUID) (*[]model.Task, error) {
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	queries := dailytreaty.New(repo.db)
	qtx := queries.WithTx(tx)

	validRows, deleteErr := qtx.DeleteTaskByID(ctx, *taskID)
	if deleteErr != nil {
		return nil, deleteErr
	}

	if validRows == 0 {
		return nil, model.ErrMissingTaskID
	}

	convertTask, convertErr := parseDBTypeToModel(ctx, qtx)
	if convertErr != nil {
		return nil, convertErr
	}

	return convertTask, tx.Commit(ctx)
}

// DeleteTaskByCategory implements [model.TaskRepository].
func (repo *TaskRepository) DeleteTaskByCategory(ctx context.Context, categoryType *string) (*[]model.Task, error) {
	tx, error := repo.db.Begin(ctx)
	if error != nil {
		return nil, error
	}

	query := dailytreaty.New(repo.db)
	qtx := query.WithTx(tx)
	validRow, deleteErr := query.DeleteTaskByCategory(ctx, *categoryType)
	if deleteErr != nil {
		return nil, deleteErr
	}

	if validRow == 0 {
		return nil, model.ErrCategoryType
	}

	convertTask, convertErr := parseDBTypeToModel(ctx, qtx)
	if convertErr != nil {
		return nil, convertErr
	}

	return convertTask, tx.Commit(ctx)
}
