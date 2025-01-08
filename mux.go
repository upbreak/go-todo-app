package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/upbreak/go-todo-app/clock"
	"github.com/upbreak/go-todo-app/config"
	"github.com/upbreak/go-todo-app/handler"
	"github.com/upbreak/go-todo-app/service"
	"github.com/upbreak/go-todo-app/store"
	"net/http"
)

func NewMux(ctx context.Context, cfg *config.DBConfig) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	r := store.Repository{Clocker: clock.RealClocker{}}
	v := validator.New()

	at := &handler.AddTask{Service: &service.AddTask{DB: db, Repo: &r}, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Service: &service.ListTask{DB: db, Repo: &r}}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux, cleanup, nil
}
