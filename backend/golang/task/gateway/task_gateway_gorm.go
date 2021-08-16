package gateway

import (
	"clean-architecture/domain/task"
	"clean-architecture/infrastructure/database"
	"context"
)

type taskGatewayGorm struct {
	*database.Gorm
}

func NewTaskGatewayGorm() task.Repository {
	gorm := database.GormConnect()
	return &taskGatewayGorm{gorm}
}

func (g *taskGatewayGorm) Save(ctx context.Context, task *task.Task) (*task.Task, error) {
	result := g.Table(TASK_TABLE_NAME).Create(TaskEntity{}.Of(task)).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (g *taskGatewayGorm) FindById(ctx context.Context, id task.Id) (*task.Task, error) {
	entity := new(TaskEntity)
	entity.Id = id.Value()

	result := g.Table(TASK_TABLE_NAME).Find(entity).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToDomain(), nil
}

func (g *taskGatewayGorm) FindAllByUserId(ctx context.Context, userId task.UserId) (*task.TaskList, error) {
	entities := []*TaskEntity{}

	result := g.Table(TASK_TABLE_NAME).Where("creator_id = ?", userId.Value()).Find(&entities).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	domains := make([]*task.Task, 0)

	for _, entity := range entities {
		domains = append(domains, entity.ToDomain())
	}

	return task.NewList(domains), nil
}

func (g *taskGatewayGorm) Update(ctx context.Context, task *task.Task) (*task.Task, error) {
	entity := TaskEntity{}.Of(task)

	result := g.Table(TASK_TABLE_NAME).Save(entity).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (g *taskGatewayGorm) Delete(ctx context.Context, id task.Id) error {
	entity := new(TaskEntity)
	entity.Id = id.Value()

	result := g.Table(TASK_TABLE_NAME).Delete(entity).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
