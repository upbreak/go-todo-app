package store

import (
	"errors"
	"github.com/upbreak/go-todo-app/entity"
)

type TaskStore struct {
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

var (
	Tasks       = &TaskStore{Tasks: make(map[entity.TaskID]*entity.Task)}
	ErrNotFound = errors.New("task not found")
)

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.IDX = ts.LastID
	ts.Tasks[t.IDX] = t
	return t.IDX, nil
}

func (ts *TaskStore) Get(t *entity.Task) (*entity.Task, error) {
	if t, ok := ts.Tasks[t.IDX]; ok {
		return t, nil
	}
	return nil, ErrNotFound
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for idx, t := range ts.Tasks {
		tasks[idx-1] = t
	}
	return tasks
}
