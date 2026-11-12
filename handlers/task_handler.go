package handlers

import (
	"net/http"

	"RNTreaty/model"

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
	// ctx := context.Context(context.Background())

	// payload := &model.TaskPayload{
	// 	CategoryType: r.Body
	// 	LastUpdateBy:
	// 	Price       :
	// }

	// taskHandler.taskUsecase.CreateTask(ctx, payload)
}
