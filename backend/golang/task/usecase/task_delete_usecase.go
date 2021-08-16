package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Delete(ctx context.Context, input task.DeleteUsecaseInput) (output *task.DeleteUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	err = u.taskRepository.Delete(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	output = new(task.DeleteUsecaseOutput)
	output.Task = result
	return

}
