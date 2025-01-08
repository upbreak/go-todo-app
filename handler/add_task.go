package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/service"
	"net/http"
)

type AddTask struct {
	Service   service.AddTasksService
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	var b entity.Task

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Println(r.Body)
		fmt.Println("handler add_task.go ServeHTTP NewDecoder error")
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
		fmt.Println("handler add_task.go ServeHTTP Validator error")
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusBadRequest)
		return
	}

	err := at.Service.AddTask(ctx, &b)

	if err != nil {
		fmt.Println("handler add_task.go ServeHTTP Repo error")
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
		Result string `json:"result"`
	}{Result: "success"}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
