package task

import "errors"

type Status struct {
	value string
}

func NewStatus(value string) (status *Status, err error) {
	switch value {
	case "INCOMPLETE", "COMPLETE":
		status = new(Status)
		status.value = value
		return
	default:
		err = errors.New("not status")
		return
	}
}

func (Status) SetInComprete() (status *Status) {
	status = new(Status)
	status.value = "INCOMPLETE"
	return
}

func (Status) SetComprete() (status *Status) {
	status = new(Status)
	status.value = "COMPLETE"
	return
}
