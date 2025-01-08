package store

import (
	"context"
	"github.com/upbreak/go-todo-app/entity"
)

type ListTasksStore interface {
	ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error)
}

type AddTaskStore interface {
	AddTask(ctx context.Context, db Beginner, t *entity.Task) error
}
