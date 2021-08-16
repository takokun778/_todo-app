package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) GetAllByUserId(ctx context.Context, input task.GetAllByUserIdUsecaseInput) (output *task.GetAllByUserIdUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindAllByUserId(timeOutContext, input.UserId)
	if err != nil {
		return nil, err
	}

	output = new(task.GetAllByUserIdUsecaseOutput)
	output.TaskList = result
	return
}
