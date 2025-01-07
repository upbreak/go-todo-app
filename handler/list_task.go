package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
	//"io"
	"net/http"
	"time"
)

type ListTask struct {
	DB   *sqlx.DB
	Repo *store.Repository
}

type task struct {
	ID    entity.TaskID `json:"id"`
	Sno   int64         `json:"sno"`
	Title string        `json:"title"`
	//Content godror.Lob    `json:"content"`
	Content string    `json:"content"`
	ShowYn  string    `json:"show_yn"`
	IsUse   string    `json:"is_use"`
	RegUno  int64     `json:"reg_uno"`
	RegUser string    `json:"reg_user"`
	RegDate time.Time `json:"reg_date"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		fmt.Println("handler list_task.go ServeHTTP error")
		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	var rsp []task
	for _, t := range tasks {

		//if t.CONTENT.Vaild {
		//	clobByte, err := io.ReadAll(t.CONTENT)
		//	if err != nil {
		//		RespondJSON(ctx, w, &ErrResponse{Message: err.Error()}, http.StatusInternalServerError)
		//	}
		//	t.CONTENT = clobByte
		//}
		rsp = append(rsp, task{
			ID:      t.IDX,
			Sno:     t.SNO,
			Title:   t.TITLE,
			Content: t.CONTENT,
			ShowYn:  t.SHOW_YN,
			IsUse:   t.IS_USE,
			RegUno:  t.REG_UNO,
			RegUser: t.REG_USER,
			RegDate: t.REG_DATE,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
