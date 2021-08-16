package task

import (
	"github.com/google/uuid"
)

type Id struct {
	value string
}

func (i *Id) Value() string {
	return i.value
}

func NewId(value string) (id *Id) {
	id = new(Id)
	id.value = value
	return
}

func CreateNewId() (id *Id) {
	id = new(Id)
	uuid, _ := uuid.NewRandom()
	id.value = uuid.String()
	return
}
