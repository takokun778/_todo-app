package task

import (
	"time"
)

type Deadline struct {
	value time.Time
}

func (t *Deadline) String() string {
	return t.value.String()
}

func NewDeadline(value time.Time) (deadline *Deadline) {
	deadline = new(Deadline)
	deadline.value = value.UTC()
	return
}

func (Deadline) Undefined() (deadline *Deadline) {
	layout := "2006-01-02 03:04:05 +0000 UTC"
	undefined := "0001-01-01 00:00:00 +0000 UTC"
	date, _ := time.Parse(layout, undefined)

	deadline = new(Deadline)
	deadline.value = date.UTC()
	return
}

func (deadline *Deadline) IsUndefined() bool {
	return deadline.value == Deadline{}.Undefined().value
}
