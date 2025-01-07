package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
	"net/http"
	"time"
)

type AddTask struct {
	DB        *sqlx.DB
	Repo      *store.Repository
	Validator *validator.Validate
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Sno   int64  `json:"sno"`
		Title string `json:"title" validate:"required"`
		//Content godror.Lob `json:"content"`
		Content string    `json:"content"`
		ShowYn  string    `json:"show_yn"`
		IsUse   string    `json:"is_use"`
		RegUno  int64     `json:"reg_uno"`
		RegUser string    `json:"reg_user"`
		RegDate time.Time `json:"reg_date"`
	}
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

	t := &entity.Task{
		SNO:      b.Sno,
		TITLE:    b.Title,
		CONTENT:  b.Content,
		SHOW_YN:  b.ShowYn,
		IS_USE:   b.IsUse,
		REG_UNO:  b.RegUno,
		REG_USER: b.RegUser,
		REG_DATE: b.RegDate,
	}

	err := at.Repo.AddTask(ctx, at.DB, t)

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
