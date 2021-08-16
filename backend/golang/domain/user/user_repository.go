package user

import "context"

type Repository interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindById(ctx context.Context, id Id) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id Id) error
}
