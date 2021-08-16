package task

import "context"

type Usecase interface {
	Create(ctx context.Context, input CreateUsecaseInput) (output *CreateUsecaseOutput, err error)
	Get(ctx context.Context, input GetUsecaseInput) (output *GetUsecaseOutput, err error)
	GetAllByUserId(ctx context.Context, input GetAllByUserIdUsecaseInput) (*GetAllByUserIdUsecaseOutput, error)
	Update(ctx context.Context, input UpdateUsecaseInput) (output *UpdateUsecaseOutput, err error)
	Done(ctx context.Context, input DoneUsecaseInput) (output *DoneUsecaseOutput, err error)
	Undone(ctx context.Context, input UndoneUsecaseInput) (output *UndoneUsecaseOutput, err error)
	Delete(ctx context.Context, input DeleteUsecaseInput) (output *DeleteUsecaseOutput, err error)
}

type CreateUsecaseInput struct {
	CreatorId CreatorId
	Title     Title
	Detail    Detail
	Deadline  Deadline
}

type CreateUsecaseOutput struct {
	Task *Task
}

type GetUsecaseInput struct {
	Id Id
}

type GetUsecaseOutput struct {
	Task *Task
}

type GetAllByUserIdUsecaseInput struct {
	UserId UserId
}

type GetAllByUserIdUsecaseOutput struct {
	TaskList *TaskList
}

type UpdateUsecaseInput struct {
	Id       Id
	Title    Title
	Detail   Detail
	Deadline Deadline
}

type UpdateUsecaseOutput struct {
	Task *Task
}

type DoneUsecaseInput struct {
	Id Id
}

type DoneUsecaseOutput struct {
	Task *Task
}

type UndoneUsecaseInput struct {
	Id Id
}

type UndoneUsecaseOutput struct {
	Task *Task
}

type DeleteUsecaseInput struct {
	Id Id
}

type DeleteUsecaseOutput struct {
	Task *Task
}
