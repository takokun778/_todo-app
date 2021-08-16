package task

type TaskList struct {
	items []*Task
}

func (t *TaskList) Items() []*Task {
	return t.items
}

func NewList(items []*Task) (taskList *TaskList) {
	taskList = new(TaskList)
	taskList.items = items
	return
}
