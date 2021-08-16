package user

import "time"

type Birthday struct {
	value time.Time
}

const BirthdayLayout = "2006-01-02"

func (b *Birthday) String() string {
	return b.value.Format(BirthdayLayout)
}

func NewBirthday(value time.Time) (birthday *Birthday) {
	birthday = new(Birthday)
	birthday.value = value
	return
}

func (Birthday) Undefined() (birthday *Birthday) {
	layout := "2006-01-02 03:04:05 +0000 UTC"
	undefined := "0001-01-01 00:00:00 +0000 UTC"
	date, _ := time.Parse(layout, undefined)

	birthday = new(Birthday)
	birthday.value = date
	return
}

func (b *Birthday) IsUndefined() bool {
	return b.value == Birthday{}.Undefined().value
}
