package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/upbreak/go-todo-app/auth"
	"github.com/upbreak/go-todo-app/clock"
	"github.com/upbreak/go-todo-app/config"
	"github.com/upbreak/go-todo-app/handler"
	"github.com/upbreak/go-todo-app/service"
	"github.com/upbreak/go-todo-app/store"
	"net/http"
)

func NewMux(ctx context.Context, cfg *config.DBConfigs) (http.Handler, []func(), error) {
	mux := chi.NewRouter()

	// 테스트용 라우팅
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	// db연결 설정
	var cleanup []func()
	safeDb, safeCleanup, err := store.New(ctx, cfg.Safe)
	if err != nil {
		cleanup = append(cleanup, func() { safeCleanup() })
		return nil, cleanup, err
	}
	_, hterpCleanup, err := store.New(ctx, cfg.Safe)
	if err != nil {
		cleanup = append(cleanup, func() { hterpCleanup() })
		return nil, cleanup, err
	}

	// jwt struct 생성
	jwt, err := auth.JwtNew(clock.RealClocker{})
	if err != nil {
		return nil, cleanup, err
	}

	r := store.Repository{Clocker: clock.RealClocker{}}
	v := validator.New()

	// 라우팅
	loginHandler := &handler.GetUser{
		Service: service.GetUser{
			DB:   safeDb,
			Repo: &r,
			Jwt:  jwt,
		},
	}
	mux.Post("/login", loginHandler.ServeHTTP)

	at := &handler.AddTask{Service: &service.AddTask{DB: safeDb, Repo: &r}, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)

	lt := &handler.ListTask{Service: &service.ListTask{DB: safeDb, Repo: &r}}
	mux.Get("/tasks", lt.ServeHTTP)

	dt := &handler.DetailTask{Service: &service.DetailTask{DB: safeDb, Repo: &r}}
	mux.Get("/tasks/{idx}", dt.ServeHTTP)

	return mux, cleanup, nil
}
