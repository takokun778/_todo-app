package gateway

import (
	"clean-architecture/domain/user"
	"clean-architecture/infrastructure/logger"
	"context"
	"strconv"
)

var users = []*user.User{}

type userGatewayInMemory struct {
}

func NewUserGatewayInMememory() user.Repository {
	return &userGatewayInMemory{}
}

func (g *userGatewayInMemory) Save(ctx context.Context, user *user.User) (*user.User, error) {
	users = append(users, user)
	for _, user := range users {
		logger.Debug(user.Id() + user.LastName() + user.FirstName() + user.Birthday().String() + strconv.FormatInt(int64(user.Age()), 10))
	}
	return user, nil
}

func (g *userGatewayInMemory) FindById(ctx context.Context, id user.Id) (*user.User, error) {
	user := make([]*user.User, 0)
	for _, value := range users {
		if id.Value() == value.Id() {
			user = append(user, value)
		}
	}
	return user[0], nil
}

func (g *userGatewayInMemory) Update(ctx context.Context, user *user.User) (*user.User, error) {
	g.delete(user)
	return g.Save(ctx, user)
}

func (g *userGatewayInMemory) Delete(ctx context.Context, id user.Id) error {
	result := []*user.User{}
	for _, value := range users {
		if id.Value() != value.Id() {
			result = append(result, value)
		}
	}
	users = result
	return nil
}

func (g *userGatewayInMemory) delete(del *user.User) error {
	result := []*user.User{}
	for _, value := range users {
		if del.Id() != value.Id() {
			result = append(result, value)
		}
	}
	users = result
	return nil
}
