package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
	"net/http"
	"time"
)

type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError)
		return
	}

	if err := at.Validator.Struct(b); err != nil {
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusBadRequest)
		return
	}

	t := &entity.Task{
		Title:   b.Title,
		Status:  "todo",
		Created: time.Now(),
	}

	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID int `json:"id"`
	}{ID: int(id)}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
