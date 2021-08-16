package task

import "time"

type Task struct {
	id        Id
	creatorId CreatorId
	title     Title
	detail    Detail
	status    Status
	createdAt Time
	updatedAt Time
	deadline  Deadline
}

func (t *Task) Id() string {
	return t.id.value
}

func (t *Task) CreatorId() string {
	return t.creatorId.value
}

func (t *Task) Title() string {
	return t.title.value
}

func (t *Task) Detail() string {
	return t.detail.value
}

func (t *Task) Status() string {
	return t.status.value
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt.value
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt.value
}

func (t *Task) Deadline() time.Time {
	return t.deadline.value
}

func (t *Task) DeadlineIsUndefined() bool {
	return t.deadline.IsUndefined()
}

type Init struct {
	Id        Id
	CreatorId CreatorId
	Title     Title
	Detail    Detail
	Status    Status
	CreatedAt Time
	UpdatedAt Time
	Deadline  Deadline
}

func New(init Init) (task *Task) {
	task = new(Task)
	task.id = init.Id
	task.creatorId = init.CreatorId
	task.title = init.Title
	task.detail = init.Detail
	task.status = init.Status
	task.createdAt = init.CreatedAt
	task.updatedAt = init.UpdatedAt
	if !init.Deadline.IsUndefined() {
		task.deadline = init.Deadline
	}
	return
}

type CreateProps struct {
	CreatorId CreatorId
	Title     Title
	Detail    Detail
	Deadline  Deadline
}

func CreateNew(props CreateProps) (task *Task) {
	now := *CreateNewTime()
	init := Init{
		Id:        *CreateNewId(),
		CreatorId: props.CreatorId,
		Title:     props.Title,
		Detail:    props.Detail,
		Status:    *Status{}.SetInComprete(),
		CreatedAt: now,
		UpdatedAt: now,
		Deadline:  props.Deadline,
	}
	task = New(init)
	return
}

type UpdateProps struct {
	Title    Title
	Detail   Detail
	Deadline Deadline
}

func (t *Task) Update(props UpdateProps) {
	if len(props.Title.value) != 0 {
		t.title = props.Title
	}
	if len(props.Detail.value) != 0 {
		t.detail = props.Detail
	}
	if !props.Deadline.IsUndefined() {
		t.deadline = props.Deadline
	}
	t.updatedAt = *CreateNewTime()
}

func (t *Task) SetInComprete() {
	t.status = *Status{}.SetInComprete()
	t.updatedAt = *CreateNewTime()
}

func (t *Task) SetComprete() {
	t.status = *Status{}.SetComprete()
	t.updatedAt = *CreateNewTime()
}
