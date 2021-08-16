package task

type Title struct {
	value string
}

func NewTitle(value string) (title *Title) {
	title = new(Title)
	title.value = value
	return
}
