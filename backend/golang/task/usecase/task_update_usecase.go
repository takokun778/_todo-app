package usecase

import (
	"clean-architecture/domain/task"
	"context"
	"time"
)

func (u *taskUsecase) Update(ctx context.Context, input task.UpdateUsecaseInput) (output *task.UpdateUsecaseOutput, err error) {
	timeOutContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := u.taskRepository.FindById(timeOutContext, input.Id)
	if err != nil {
		return nil, err
	}

	props := task.UpdateProps{
		Title:    input.Title,
		Detail:   input.Detail,
		Deadline: input.Deadline,
	}

	result.Update(props)

	result, err = u.taskRepository.Update(timeOutContext, result)

	output = new(task.UpdateUsecaseOutput)
	output.Task = result
	return
}
