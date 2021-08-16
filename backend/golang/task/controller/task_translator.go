package controller

import (
	"clean-architecture/domain/task"
	pbTask "clean-architecture/infrastructure/proto/task"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskTranslator struct{}

func (TaskTranslator) ToDomainId(value string) (*task.Id, error) {
	if value == "" {
		return nil, errors.New("task id is empty")
	}
	return task.NewId(value), nil
}

func (TaskTranslator) ToDomainCreatorId(value string) (*task.CreatorId, error) {
	if value == "" {
		return nil, errors.New("creator id is empty")
	}
	return task.NewCreatorId(value), nil
}

func (TaskTranslator) ToDomainTitle(value string) (*task.Title, error) {
	if value == "" {
		return nil, errors.New("title is empty")
	}
	return task.NewTitle(value), nil
}

func (TaskTranslator) ToDomainDetail(value string) *task.Detail {
	return task.NewDetail(value)
}

func (TaskTranslator) ToDomainDeadline(value *timestamppb.Timestamp) *task.Deadline {
	if value == nil {
		return task.Deadline{}.Undefined()
	}
	return task.NewDeadline(value.AsTime().UTC())
}

func (TaskTranslator) ToDomainUserId(value string) (*task.UserId, error) {
	if value == "" {
		return nil, errors.New("user id is empty")
	}
	return task.NewUserId(value), nil
}

func (TaskTranslator) ToProto(domain *task.Task) (proto *pbTask.Task) {
	proto = new(pbTask.Task)
	proto.Id = domain.Id()
	proto.CreatorId = domain.CreatorId()
	proto.Title = domain.Title()
	proto.Detail = domain.Detail()
	switch domain.Status() {
	case "INCOMPLETE":
		proto.Status = 0
	case "COMPLETE":
		proto.Status = 1
	}
	proto.CreatedAt = timestamppb.New(domain.CreatedAt())
	proto.UpdatedAt = timestamppb.New(domain.UpdatedAt())
	if !domain.DeadlineIsUndefined() {
		proto.Deadline = timestamppb.New(domain.Deadline())
	}
	return
}

func (TaskTranslator) ToProtoList(domain []*task.Task) (proto []*pbTask.Task) {
	proto = make([]*pbTask.Task, 0)
	for _, value := range domain {
		proto = append(proto, TaskTranslator{}.ToProto(value))
	}
	return proto
}
