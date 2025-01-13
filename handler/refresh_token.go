package handler

import (
	"encoding/json"
	"github.com/upbreak/go-todo-app/auth"
	"net/http"
)

type RefreshToken struct {
	Jwt *auth.JWTUtils
}

func (rt *RefreshToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var token struct {
		RefreshToken string `json:"refreshToken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError)
		return
	}

	jwtClaims, err := rt.Jwt.RefreshToken(token.RefreshToken)
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
		Result       string `json:"Result"`
		Token        string `json:"Token"`
		RefreshToken string `json:"RefreshToken"`
	}{Result: "Success", Token: jwtClaims.Token, RefreshToken: jwtClaims.RefreshToken}

	RespondJSON(ctx, w, &rsp, http.StatusOK)
}
