package usecase

import (
	"clean-architecture/domain/task"
)

type taskUsecase struct {
	taskRepository task.Repository
}

func NewTaskUsecase(taskRepository task.Repository) task.Usecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}
