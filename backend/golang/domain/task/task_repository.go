package task

import "context"

type Repository interface {
	Save(ctx context.Context, task *Task) (*Task, error)
	FindById(ctx context.Context, id Id) (*Task, error)
	FindAllByUserId(ctx context.Context, userId UserId) (*TaskList, error)
	Update(ctx context.Context, task *Task) (*Task, error)
	Delete(ctx context.Context, id Id) error
}
