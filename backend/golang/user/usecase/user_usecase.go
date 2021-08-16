package usecase

import (
	"clean-architecture/domain/user"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUserUsecase(userRepository user.Repository) user.Usecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
