package handler

import (
	"github.com/upbreak/go-todo-app/auth"
	"net/http"
)

// api호출시 jwt를 확인하는 미들웨어
func AuthMiddleware(jwt *auth.JWTUtils) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req, err := jwt.FillContext(r)
			if err != nil {
				RespondJSON(
					r.Context(),
					w,
					ErrResponse{Message: "not found auth info", Details: []string{err.Error()}},
					http.StatusUnauthorized,
				)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
