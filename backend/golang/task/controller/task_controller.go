package controller

import (
	"clean-architecture/domain/task"
	pbTask "clean-architecture/infrastructure/proto/task"
	"context"
)

type taskController struct {
	taskUsecase task.Usecase
	pbTask.UnimplementedTaskServiceServer
}

func NewtaskController(taskUsecase task.Usecase) *taskController {
	return &taskController{
		taskUsecase: taskUsecase,
	}
}

func (c *taskController) Create(ctx context.Context, req *pbTask.CreateRequest) (res *pbTask.CreateResponse, err error) {
	creatorId, err := TaskTranslator{}.ToDomainCreatorId(req.CreatorId)
	if err != nil {
		return nil, err
	}

	title, err := TaskTranslator{}.ToDomainTitle(req.Title)
	if err != nil {
		return nil, err
	}

	input := task.CreateUsecaseInput{
		CreatorId: *creatorId,
		Title:     *title,
		Detail:    *TaskTranslator{}.ToDomainDetail(req.Detail),
		Deadline:  *TaskTranslator{}.ToDomainDeadline(req.Deadline),
	}

	output, err := c.taskUsecase.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.CreateResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}

func (c *taskController) Get(ctx context.Context, req *pbTask.GetRequest) (res *pbTask.GetResponse, err error) {
	id, err := TaskTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := task.GetUsecaseInput{
		Id: *id,
	}

	output, err := c.taskUsecase.Get(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.GetResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}

func (c *taskController) List(ctx context.Context, req *pbTask.ListRequest) (res *pbTask.ListResponse, err error) {
	userId, err := TaskTranslator{}.ToDomainUserId(req.UserId)
	if err != nil {
		return nil, err
	}

	input := task.GetAllByUserIdUsecaseInput{
		UserId: *userId,
	}

	output, err := c.taskUsecase.GetAllByUserId(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.ListResponse)
	res.Task = TaskTranslator{}.ToProtoList(output.TaskList.Items())
	return
}

func (c *taskController) Update(ctx context.Context, req *pbTask.UpdateRequest) (res *pbTask.UpdateResponse, err error) {
	id, err := TaskTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	title, err := TaskTranslator{}.ToDomainTitle(req.Title)
	if err != nil {
		title, _ = TaskTranslator{}.ToDomainTitle("")
	}

	input := task.UpdateUsecaseInput{
		Id:       *id,
		Title:    *title,
		Detail:   *TaskTranslator{}.ToDomainDetail(req.Detail),
		Deadline: *TaskTranslator{}.ToDomainDeadline(req.Deadline),
	}

	output, err := c.taskUsecase.Update(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.UpdateResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}

func (c *taskController) Done(ctx context.Context, req *pbTask.DoneRequest) (res *pbTask.DoneResponse, err error) {
	id, err := TaskTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := task.DoneUsecaseInput{
		Id: *id,
	}

	output, err := c.taskUsecase.Done(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.DoneResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}

func (c *taskController) Undone(ctx context.Context, req *pbTask.UndoneRequest) (res *pbTask.UndoneResponse, err error) {
	id, err := TaskTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := task.UndoneUsecaseInput{
		Id: *id,
	}

	output, err := c.taskUsecase.Undone(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.UndoneResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}

func (c *taskController) Delete(ctx context.Context, req *pbTask.DeleteRequest) (res *pbTask.DeleteResponse, err error) {
	id, err := TaskTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := task.DeleteUsecaseInput{
		Id: *id,
	}

	output, err := c.taskUsecase.Delete(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbTask.DeleteResponse)
	res.Task = TaskTranslator{}.ToProto(output.Task)
	return
}
