package usecase

import (
	"clean-architecture/domain/user"
	"context"
	"time"
)

func (u *userUsecase) Get(ctx context.Context, input user.GetUsecaseInput) (output *user.GetUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.userRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	output = new(user.GetUsecaseOutput)
	output.User = result
	return
}
