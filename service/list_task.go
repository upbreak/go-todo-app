package service

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo store.ListTasksStore
}

func (l *ListTask) ListTasks(ctx context.Context) ([]entity.TaskListRsp, error) {
	tasks, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, fmt.Errorf("fail to list task: %w", err)
	}

	var rsp []entity.TaskListRsp
	for _, t := range tasks {
		rsp = append(rsp, entity.TaskListRsp{
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

	return rsp, nil
}
