package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Get(ctx context.Context, input task.GetUsecaseInput) (output *task.GetUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	output = new(task.GetUsecaseOutput)
	output.Task = result
	return
}
