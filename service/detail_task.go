package service

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
)

type DetailTask struct {
	DB   store.Queryer
	Repo store.DetailTaskStore
}

func (d *DetailTask) DetailTask(ctx context.Context, idx string) (entity.TaskListRsp, error) {
	task, err := d.Repo.DetailTask(ctx, d.DB, idx)

	if err != nil {
		return entity.TaskListRsp{}, fmt.Errorf("fail to detail task: %w", err)
	}

	rsp := entity.TaskListRsp{
		ID:      task.IDX,
		Sno:     task.SNO,
		Title:   task.TITLE,
		Content: task.CONTENT,
		ShowYn:  task.SHOW_YN,
		IsUse:   task.IS_USE,
		RegUno:  task.REG_UNO,
		RegUser: task.REG_USER,
		RegDate: task.REG_DATE,
	}

	return rsp, nil

}
