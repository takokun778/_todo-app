package usecase

import (
	"clean-architecture/domain/user"
	"context"
	"time"
)

func (u *userUsecase) Register(ctx context.Context, input user.RegisterUsecaseInput) (output *user.RegisterUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	props := user.CreateProps(input)

	result := user.CreateNew(props)

	result, err = u.userRepository.Save(timeOutContext, result)
	if err != nil {
		return nil, err
	}

	output = new(user.RegisterUsecaseOutput)
	output.User = result
	return
}
