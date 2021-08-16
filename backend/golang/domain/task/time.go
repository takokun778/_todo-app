package task

import (
	"time"
)

type Time struct {
	value time.Time
}

func (t *Time) String() string {
	return t.value.String()
}

func NewTime(value time.Time) (time *Time) {
	time = new(Time)
	time.value = value.UTC()
	return
}

func CreateNewTime() (result *Time) {
	result = new(Time)
	result.value = time.Now().UTC()
	return
}
