package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"RNTreaty/model"
	"RNTreaty/utils"

	"github.com/go-chi/chi"
)

type TaskHandler struct {
	taskUsecase model.TaskUsecase
}

func NewTaskHandler(taskUsecase model.TaskUsecase) *chi.Mux {
	taskHandler := &TaskHandler{
		taskUsecase,
	}

	route := chi.NewRouter()

	route.Post("/create", taskHandler.createTask)
	route.Delete("/delete", taskHandler.deleteTask)
	return route
}

func (taskHandler *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	ctx := context.Context(context.Background())

	payload := &model.TaskPayload{}
	if decodeErr := utils.DecodeJSON(r, payload); decodeErr != nil {
		slog.Error(decodeErr.Error())
	}

	if taskErr := taskHandler.taskUsecase.CreateTask(ctx, payload); taskErr != nil {
		slog.Error(taskErr.Error())
	}

	if writeErr := utils.WriteJSON(w, http.StatusOK, payload); writeErr != nil {
		slog.Error(writeErr.Error())
	}

	slog.Info(fmt.Sprintf("Created Task API requested by %s", payload.LastUpdateBy))
}

func (taskHandler *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payload := &model.TaskDeletePayload{}
	if decodeErr := utils.DecodeJSON(r, payload); decodeErr != nil {
		slog.Error(decodeErr.Error())
	}

	var allTask *[]model.Task
	var taskErr error
	if payload.Taskid != "" || payload.TaskCategory != "" {
		allTask, taskErr = taskHandler.taskUsecase.DeleteTaskByID(ctx, &payload.Taskid)
		slog.Info(fmt.Sprintf("Delete Task by ID API requested by %s", payload.LastUpdateBy))
	}

	if payload.Taskid == "" && payload.TaskCategory != "" {
		allTask, taskErr = taskHandler.taskUsecase.DeleteTaskByCategory(ctx, &payload.TaskCategory)
		slog.Info(fmt.Sprintf("Delete Task by Category API requested by %s", payload.LastUpdateBy))
	}

	deleteResponse := struct {
		Response *model.GenericResponse `json:"response"`
		TaskList *[]model.Task          `json:"tasklist"`
	}{}
	if taskErr != nil {
		slog.Error(taskErr.Error())
		deleteFailedResponse := &model.GenericResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Failed to deleted task %s", taskErr.Error()),
			Success: false,
		}
		deleteResponse.Response = deleteFailedResponse
	} else {
		deleteSuccessResponse := &model.GenericResponse{
			Code:    http.StatusOK,
			Message: "Successfully deleted task",
			Success: true,
		}
		deleteResponse.Response = deleteSuccessResponse
	}

	deleteResponse.TaskList = allTask
	if writeErr := utils.WriteJSON(w, deleteResponse.Response.Code, deleteResponse); writeErr != nil {
		slog.Error(writeErr.Error())
	}
}
