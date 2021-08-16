package user

import "context"

type Usecase interface {
	Register(ctx context.Context, input RegisterUsecaseInput) (output *RegisterUsecaseOutput, err error)
	Get(ctx context.Context, input GetUsecaseInput) (output *GetUsecaseOutput, err error)
	Update(ctx context.Context, input UpdateUsecaseInput) (output *UpdateUsecaseOutput, err error)
	Terminate(ctx context.Context, input TerminateUsecaseInput) (output *TerminateUsecaseOutput, err error)
}

type RegisterUsecaseInput struct {
	FirstName Name
	LastName  Name
	Birthday  Birthday
}

type RegisterUsecaseOutput struct {
	User *User
}

type GetUsecaseInput struct {
	Id Id
}

type GetUsecaseOutput struct {
	User *User
}

type UpdateUsecaseInput struct {
	Id        Id
	FirstName Name
	LastName  Name
	Birthday  Birthday
}

type UpdateUsecaseOutput struct {
	User *User
}

type TerminateUsecaseInput struct {
	Id Id
}

type TerminateUsecaseOutput struct {
	User *User
}
