package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/upbreak/go-todo-app/service"
	"net/http"
)

type DetailTask struct {
	Service service.DetailTaskService
}

func (d *DetailTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idx := chi.URLParam(r, "idx")

	rsp, err := d.Service.DetailTask(ctx, idx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	RespondJSON(ctx, w, rsp, http.StatusOK)

}
