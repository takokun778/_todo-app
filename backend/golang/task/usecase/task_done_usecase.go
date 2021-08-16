package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Done(ctx context.Context, input task.DoneUsecaseInput) (output *task.DoneUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	result.SetComprete()

	result, err = u.taskRepository.Update(timeOutContext, result)

	output = new(task.DoneUsecaseOutput)
	output.Task = result
	return
}
