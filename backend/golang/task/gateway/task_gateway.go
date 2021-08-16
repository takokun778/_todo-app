package gateway

import (
	"clean-architecture/domain/task"
	"context"
)

type taskGateway struct {
}

func NewTaskGateway() task.Repository {
	return &taskGateway{}
}

func (g *taskGateway) Save(ctx context.Context, task *task.Task) (*task.Task, error) {
	return task, nil
}

func (g *taskGateway) FindById(ctx context.Context, id task.Id) (*task.Task, error) {
	return nil, nil
}

func (g *taskGateway) FindAllByUserId(ctx context.Context, userId task.UserId) (*task.TaskList, error) {
	return nil, nil
}

func (g *taskGateway) Update(ctx context.Context, task *task.Task) (*task.Task, error) {
	return task, nil
}

func (g *taskGateway) Delete(ctx context.Context, id task.Id) error {
	return nil
}
