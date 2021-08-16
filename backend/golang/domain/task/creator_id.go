package task

type CreatorId struct {
	value string
}

func NewCreatorId(value string) (creatorId *CreatorId) {
	creatorId = new(CreatorId)
	creatorId.value = value
	return
}
