package handler

import (
	"encoding/json"
	"github.com/upbreak/go-todo-app/service"
	"net/http"
)

type GetUser struct {
	Service service.GetUser
}

func (g *GetUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var login struct {
		UserId string `json:"userId"`
		Pw     string `json:"pw"`
	}

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		RespondJSON(
			ctx,
			w,
			&ErrResponse{
				Message: err.Error(),
			},
			http.StatusInternalServerError)
		return
	}

	jwtClaims, err := g.Service.GetUserValid(ctx, login.UserId, login.Pw)
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
		Result string `json:"Result"`
		Token  string `json:"Token"`
	}{Result: "Success", Token: jwtClaims.Token}

	RespondJSON(ctx, w, &rsp, http.StatusOK)
}
