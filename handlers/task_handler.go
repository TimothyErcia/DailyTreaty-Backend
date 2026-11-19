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
	route.Delete("/delete", taskHandler.deleteTask)
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

	if taskErr != nil {
		slog.Error(taskErr.Error())
		errorResponse(w, http.StatusBadRequest, fmt.Sprintf("Failed to deleted task %s", taskErr.Error()))
		return
	}

	deleteSuccessResponse := &model.GenericResponse{
		Code:    http.StatusOK,
		Message: "Successfully deleted task",
		Success: true,
	}
	deleteResponse := struct {
		Response *model.GenericResponse `json:"response"`
		TaskList *[]model.Task          `json:"tasklist"`
	}{Response: deleteSuccessResponse, TaskList: allTask}
	if writeErr := utils.WriteJSON(w, deleteResponse.Response.Code, deleteResponse); writeErr != nil {
		slog.Error(writeErr.Error())
	}
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
