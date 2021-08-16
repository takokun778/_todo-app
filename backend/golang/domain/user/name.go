package user

type Name struct {
	value string
}

func NewName(value string) (name *Name) {
	name = new(Name)
	name.value = value
	return
}
