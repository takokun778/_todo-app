package gateway

import (
	"clean-architecture/domain/user"
	"time"
)

const USER_TABLE_NAME = "users"

type UserEntity struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday" sql:"type:date"`
}

func (UserEntity) Of(domain *user.User) (entity *UserEntity) {
	entity = new(UserEntity)
	entity.Id = domain.Id()
	entity.FirstName = domain.FirstName()
	entity.LastName = domain.LastName()
	layout := "2006-01-02"
	entity.Birthday = domain.Birthday().Format(layout)
	return
}

func (e *UserEntity) ToDomain() (domain *user.User) {
	id := user.NewId(e.Id)
	firstName := user.NewName(e.FirstName)
	lastName := user.NewName(e.LastName)
	layout := "2006-01-02T03:04:00Z"
	date, _ := time.Parse(layout, e.Birthday)
	birthday := user.NewBirthday(date)

	init := user.Init{
		Id:        *id,
		FirstName: *firstName,
		LastName:  *lastName,
		Birthday:  *birthday,
	}

	domain = user.New(init)
	return
}
