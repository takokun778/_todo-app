package usecase

import (
	"clean-architecture/domain/user"
	"context"
	"time"
)

func (u *userUsecase) Terminate(ctx context.Context, input user.TerminateUsecaseInput) (output *user.TerminateUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.userRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	err = u.userRepository.Delete(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	output = new(user.TerminateUsecaseOutput)
	output.User = result
	return
}
