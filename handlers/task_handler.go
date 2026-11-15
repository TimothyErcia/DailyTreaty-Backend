package handlers

import (
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
	route.Patch("/update", taskHandler.updateTask)
	return route
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	response := &model.GenericResponse{
		Code:    status,
		Message: message,
		Success: false,
	}

	utils.WriteJSON(w, response.Code, response)
}

func (taskHandler *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payload := &model.TaskPayload{}
	if decodeErr := utils.DecodeJSON(r, payload); decodeErr != nil {
		errorResponse(w, http.StatusBadRequest, decodeErr.Error())
		slog.Error(decodeErr.Error())
	}

	createdTask, taskErr := taskHandler.taskUsecase.CreateTask(ctx, payload)
	if taskErr != nil {
		errorResponse(w, http.StatusBadRequest, taskErr.Error())
		slog.Error(taskErr.Error())
	}

	successResponse := &model.GenericResponse{
		Code:    http.StatusOK,
		Message: "Task successfully created",
		Success: true,
	}
	createResponse := struct {
		Task     *model.Task            `json:"task"`
		Response *model.GenericResponse `json:"response"`
	}{Task: createdTask, Response: successResponse}

	if writeErr := utils.WriteJSON(w, successResponse.Code, &createResponse); writeErr != nil {
		errorResponse(w, http.StatusBadRequest, writeErr.Error())
		slog.Error(writeErr.Error())
	}

	slog.Info(fmt.Sprintf("Created Task API requested by %s", payload.LastUpdateBy))
}

func (taskHandler *TaskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payload := &model.TaskUpdatePayload{}
	if decodeErr := utils.DecodeJSON(r, payload); decodeErr != nil {
		errorResponse(w, http.StatusBadRequest, decodeErr.Error())
		slog.Error(decodeErr.Error())
	}

	if updateErr := taskHandler.taskUsecase.UpdateTask(ctx, &payload.Taskid, &payload.TaskPayload); updateErr != nil {
		errorResponse(w, http.StatusBadRequest, updateErr.Error())
		slog.Error(updateErr.Error())
	}

	updateResponse := &model.GenericResponse{
		Code:    http.StatusOK,
		Message: "Task updated successfully",
		Success: true,
	}
	if writeErr := utils.WriteJSON(w, updateResponse.Code, updateResponse); writeErr != nil {
		errorResponse(w, http.StatusBadRequest, writeErr.Error())
		slog.Error(writeErr.Error())
	}

	slog.Info(fmt.Sprintf("Update Task API requested by %s", payload.LastUpdateBy))
}
