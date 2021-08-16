package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Undone(ctx context.Context, input task.UndoneUsecaseInput) (output *task.UndoneUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	result.SetInComprete()

	result, err = u.taskRepository.Update(timeOutContext, result)

	output = new(task.UndoneUsecaseOutput)
	output.Task = result
	return
}
