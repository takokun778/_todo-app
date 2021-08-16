package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Create(ctx context.Context, input task.CreateUsecaseInput) (output *task.CreateUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	props := task.CreateProps(input)

	result := task.CreateNew(props)

	result, err = u.taskRepository.Save(timeOutContext, result)
	if err != nil {
		return nil, err
	}

	output = new(task.CreateUsecaseOutput)
	output.Task = result
	return
}
