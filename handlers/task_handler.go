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
