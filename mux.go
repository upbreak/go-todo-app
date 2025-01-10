package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/rs/cors"
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

	// CORS 미들웨어 설정
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://example.com", "http://localhost:3000"}, // 허용할 도메인
		AllowCredentials: true,                                                    // 쿠키 허용
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},     // 허용할 메서드
		AllowedHeaders:   []string{"Content-Type", "Authorization"},               // 허용할 헤더
	})

	// db연결 설정
	var cleanup []func()
	safeDb, safeCleanup, err := store.New(ctx, cfg.Safe)
	cleanup = append(cleanup, func() { safeCleanup() })
	if err != nil {
		return nil, cleanup, err
	}
	_, hterpCleanup, err := store.New(ctx, cfg.Safe)
	cleanup = append(cleanup, func() { hterpCleanup() })
	if err != nil {
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
	lt := &handler.ListTask{Service: &service.ListTask{DB: safeDb, Repo: &r}}
	dt := &handler.DetailTask{Service: &service.DetailTask{DB: safeDb, Repo: &r}}

	// 미들웨어를 사용하여 토큰 검사 후 ServeHTTP 실행
	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwt))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
		r.Get("/{idx}", dt.ServeHTTP)
	})

	handlerMux := c.Handler(mux)

	return handlerMux, cleanup, nil
}
