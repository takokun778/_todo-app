package gateway

import (
	"clean-architecture/domain/user"
	"clean-architecture/infrastructure/database"
	"context"
)

type userGatewayGorm struct {
	*database.Gorm
}

func NewUserGatewayGorm() user.Repository {
	gorm := database.GormConnect()
	return &userGatewayGorm{gorm}
}

func (g *userGatewayGorm) Save(ctx context.Context, user *user.User) (*user.User, error) {
	result := g.Table(USER_TABLE_NAME).Create(UserEntity{}.Of(user)).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (g *userGatewayGorm) FindById(ctx context.Context, id user.Id) (*user.User, error) {
	entity := new(UserEntity)
	entity.Id = id.Value()

	result := g.Table(USER_TABLE_NAME).Find(entity).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToDomain(), nil
}

func (g *userGatewayGorm) Update(ctx context.Context, user *user.User) (*user.User, error) {
	entity := UserEntity{}.Of(user)

	result := g.Table(USER_TABLE_NAME).Save(entity).WithContext(ctx)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (g *userGatewayGorm) Delete(ctx context.Context, id user.Id) error {
	entity := new(UserEntity)
	entity.Id = id.Value()

	result := g.Table(USER_TABLE_NAME).Delete(entity).WithContext(ctx)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
