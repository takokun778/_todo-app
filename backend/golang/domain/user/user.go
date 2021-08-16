package user

import (
	"time"
)

type User struct {
	id        Id
	firstName Name
	lastName  Name
	birthday  Birthday
	age       Age
}

func (u *User) Id() string {
	return u.id.value
}

func (u *User) FirstName() string {
	return u.firstName.value
}

func (u *User) LastName() string {
	return u.lastName.value
}

func (u *User) Birthday() time.Time {
	return u.birthday.value
}

func (u *User) Age() int {
	return u.age.value
}

type Init struct {
	Id        Id
	FirstName Name
	LastName  Name
	Birthday  Birthday
}

func New(init Init) (user *User) {
	user = new(User)
	user.id = init.Id
	user.firstName = init.FirstName
	user.lastName = init.LastName
	user.birthday = init.Birthday
	user.age = *NewAge(init.Birthday.value)
	return
}

type CreateProps struct {
	FirstName Name
	LastName  Name
	Birthday  Birthday
}

func CreateNew(props CreateProps) (user *User) {
	init := &Init{
		Id:        *CreateNewId(),
		FirstName: props.FirstName,
		LastName:  props.LastName,
		Birthday:  props.Birthday,
	}
	user = New(*init)
	return
}

type UpdateProps struct {
	FirstName Name
	LastName  Name
	Birthday  Birthday
}

func (u *User) Update(props UpdateProps) {
	if len(props.FirstName.value) != 0 {
		u.firstName = props.FirstName
	}
	if len(props.LastName.value) != 0 {
		u.lastName = props.LastName
	}
	if !props.Birthday.IsUndefined() {
		u.birthday = props.Birthday
		u.age = *NewAge(props.Birthday.value)
	}
}
