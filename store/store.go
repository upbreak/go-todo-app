package store

import (
	"context"
	"github.com/upbreak/go-todo-app/entity"
)

type ListTasksStore interface {
	ListTasks(ctx context.Context, db Queryer) (entity.Tasks, error)
}

type DetailTaskStore interface {
	DetailTask(ctx context.Context, db Queryer, idx string) (entity.Task, error)
}

type AddTaskStore interface {
	AddTask(ctx context.Context, db Beginner, t *entity.Task) error
}

type GetUserValidStore interface {
	GetUserValid(ctx context.Context, db Queryer, id string, pwMd5 string) (entity.User, error)
}
