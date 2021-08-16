package task

type UserId struct {
	value string
}

func (u *UserId) Value() string {
	return u.value
}

func NewUserId(value string) (userId *UserId) {
	userId = new(UserId)
	userId.value = value
	return userId
}
