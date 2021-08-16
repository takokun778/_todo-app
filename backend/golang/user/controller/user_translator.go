package controller

import (
	"clean-architecture/domain/user"
	pbUser "clean-architecture/infrastructure/proto/user"
	"errors"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
)

type UserTranslator struct{}

func (UserTranslator) ToDomainId(value string) (*user.Id, error) {
	if value == "" {
		return nil, errors.New("user id is empty")
	}
	return user.NewId(value), nil
}

func (UserTranslator) ToDomainName(value string) (*user.Name, error) {
	if value == "" {
		return nil, errors.New("user name is empty")
	}
	return user.NewName(value), nil
}

func (UserTranslator) ToDomainBirthday(value date.Date) (*user.Birthday, error) {
	year := fmt.Sprintf("%04d", value.Year)
	month := fmt.Sprintf("%02d", value.Month)
	day := fmt.Sprintf("%02d", value.Day)
	parse, err := time.Parse("2006-01-02", year+"-"+month+"-"+day)
	if err != nil {
		return nil, err
	}
	return user.NewBirthday(parse), nil
}

func (UserTranslator) ToProto(domain *user.User) (proto *pbUser.User) {
	proto = new(pbUser.User)
	proto.Id = domain.Id()
	proto.FirstName = domain.FirstName()
	proto.LastName = domain.LastName()
	proto.Birthday = &date.Date{
		Year:  int32(domain.Birthday().Year()),
		Month: int32(domain.Birthday().Month()),
		Day:   int32(domain.Birthday().Day()),
	}
	proto.Age = int64(domain.Age())
	return
}
