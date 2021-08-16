package gateway

import (
	"clean-architecture/domain/user"
	"context"
)

type userGateway struct {
}

func NewUserGateway() user.Repository {
	return &userGateway{}
}

func (g *userGateway) Save(ctx context.Context, user *user.User) (*user.User, error) {
	return user, nil
}

func (g *userGateway) FindById(ctx context.Context, id user.Id) (*user.User, error) {
	return nil, nil
}

func (g *userGateway) Update(ctx context.Context, user *user.User) (*user.User, error) {
	return user, nil
}

func (g *userGateway) Delete(ctx context.Context, id user.Id) error {
	return nil
}
