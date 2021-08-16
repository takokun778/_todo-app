package user

import "time"

type Age struct {
	value int
}

func NewAge(birthday time.Time) (age *Age) {
	age = new(Age)
	age.value = time.Now().Year() - birthday.Year()
	return
}
