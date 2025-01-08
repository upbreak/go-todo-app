package service

import (
	"context"
	"fmt"
	"github.com/upbreak/go-todo-app/entity"
	"github.com/upbreak/go-todo-app/store"
)

type AddTask struct {
	DB   store.Beginner
	Repo store.AddTaskStore
}

func (a *AddTask) AddTask(ctx context.Context, t *entity.Task) error {
	if err := a.Repo.AddTask(ctx, a.DB, t); err != nil {
		return fmt.Errorf("fail to add task: %w", err)
	}
	return nil
}
