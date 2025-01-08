package handler

import (
	"fmt"
	"github.com/upbreak/go-todo-app/service"
	"net/http"
)

type ListTask struct {
	Service service.ListTasksService
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rsp, err := lt.Service.ListTasks(ctx)
	if err != nil {
		fmt.Println("handler list_task.go ServeHTTP error")
		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
