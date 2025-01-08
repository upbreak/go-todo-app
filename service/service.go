package service

import (
	"context"
	"github.com/upbreak/go-todo-app/entity"
)

type ListTasksService interface {
	ListTasks(ctx context.Context) ([]entity.TaskListRsp, error)
}

type DetailTaskService interface {
	DetailTask(ctx context.Context, idx string) (entity.TaskListRsp, error)
}

type AddTasksService interface {
	AddTask(ctx context.Context, t *entity.Task) error
}