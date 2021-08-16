package gateway

import (
	"clean-architecture/domain/task"
	"context"
	"log"
)

var tasks = []*task.Task{}

type taskGatewayInMemory struct {
}

func NewTaskGatewayInMemory() task.Repository {
	return &taskGatewayInMemory{}
}

func (g *taskGatewayInMemory) Save(ctx context.Context, task *task.Task) (*task.Task, error) {
	tasks = append(tasks, task)
	for _, task := range tasks {
		log.Println(task.Id() + " " + task.CreatorId() + " " + task.Title() + " " + task.Detail() + " " + task.Status())
		log.Println(task.CreatedAt().String() + " " + task.UpdatedAt().String() + " " + task.Deadline().String())
	}
	return task, nil
}

func (g *taskGatewayInMemory) FindById(ctx context.Context, id task.Id) (*task.Task, error) {
	task := make([]*task.Task, 0)
	for _, value := range tasks {
		if id.Value() == value.Id() {
			task = append(task, value)
		}
	}
	return task[0], nil
}

func (g *taskGatewayInMemory) FindAllByUserId(ctx context.Context, userId task.UserId) (*task.TaskList, error) {
	result := make([]*task.Task, 0)
	for _, value := range tasks {
		if userId.Value() == value.CreatorId() {
			result = append(result, value)
		}
	}
	return task.NewList(result), nil
}

func (g *taskGatewayInMemory) Update(ctx context.Context, task *task.Task) (*task.Task, error) {
	g.delete(task)
	return g.Save(ctx, task)
}

func (g *taskGatewayInMemory) Delete(ctx context.Context, id task.Id) error {
	result := []*task.Task{}
	for _, value := range tasks {
		if id.Value() != value.Id() {
			result = append(result, value)
		}
	}
	tasks = result
	return nil
}

func (g *taskGatewayInMemory) delete(del *task.Task) error {
	result := []*task.Task{}
	for _, value := range tasks {
		if del.Id() != value.Id() {
			result = append(result, value)
		}
	}
	tasks = result
	return nil
}
