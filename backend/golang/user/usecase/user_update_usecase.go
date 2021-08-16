package usecase

import (
	"clean-architecture/domain/user"
	"context"
	"time"
)

func (u *userUsecase) Update(ctx context.Context, input user.UpdateUsecaseInput) (output *user.UpdateUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.userRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	props := user.UpdateProps{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Birthday:  input.Birthday,
	}

	result.Update(props)

	result, err = u.userRepository.Update(timeOutContext, result)
	if err != nil {
		return nil, err
	}

	output = new(user.UpdateUsecaseOutput)
	output.User = result
	return
}
