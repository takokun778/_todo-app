package gateway

import (
	"clean-architecture/domain/task"
	"time"
)

const TASK_TABLE_NAME = "tasks"

type TaskEntity struct {
	Id        string     `json:"id"`
	CreatorId string     `json:"creator_id"`
	Title     string     `json:"title"`
	Detail    string     `json:"detail"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Deadline  *time.Time `json:"deadline"`
	// 時刻の型を string にしたら時刻変換でバグる
}

func (TaskEntity) Of(domain *task.Task) (entity *TaskEntity) {
	entity = new(TaskEntity)
	entity.Id = domain.Id()
	entity.CreatorId = domain.CreatorId()
	entity.Title = domain.Title()
	entity.Detail = domain.Detail()
	entity.Status = domain.Status()
	entity.CreatedAt = domain.CreatedAt().UTC()
	entity.UpdatedAt = domain.UpdatedAt().UTC()
	if domain.DeadlineIsUndefined() {
		entity.Deadline = nil
	} else {
		deadline := domain.Deadline().UTC()
		entity.Deadline = &deadline
	}
	return
}

func (e *TaskEntity) ToDomain() (domain *task.Task) {
	id := task.NewId(e.Id)
	creatorId := task.NewCreatorId(e.CreatorId)
	title := task.NewTitle(e.Title)
	detail := task.NewDetail(e.Detail)
	status, _ := task.NewStatus(e.Status)
	createdAt := task.NewTime(e.CreatedAt.UTC())
	updatedAt := task.NewTime(e.UpdatedAt.UTC())
	deadline := task.Deadline{}
	if e.Deadline == nil {
		deadline = *task.Deadline{}.Undefined()
	} else {
		deadline = *task.NewDeadline(e.Deadline.UTC())
	}

	init := task.Init{
		Id:        *id,
		CreatorId: *creatorId,
		Title:     *title,
		Detail:    *detail,
		Status:    *status,
		CreatedAt: *createdAt,
		UpdatedAt: *updatedAt,
		Deadline:  deadline,
	}

	domain = task.New(init)
	return

}
