package controller

import (
	"clean-architecture/domain/user"
	pbUser "clean-architecture/infrastructure/proto/user"
	"context"
)

type userController struct {
	userUsecase user.Usecase
	pbUser.UnimplementedUserServiceServer
}

func NewUserController(userUsecase user.Usecase) *userController {
	return &userController{
		userUsecase: userUsecase,
	}
}

func (c *userController) Register(ctx context.Context, req *pbUser.RegisterRequest) (res *pbUser.RegisterResponse, err error) {
	firstName, err := UserTranslator{}.ToDomainName(req.FirstName)
	if err != nil {
		return nil, err
	}

	lastName, err := UserTranslator{}.ToDomainName(req.LastName)
	if err != nil {
		return nil, err
	}

	birthday, err := UserTranslator{}.ToDomainBirthday(*req.Birthday)
	if err != nil {
		return nil, err
	}

	input := user.RegisterUsecaseInput{
		FirstName: *firstName,
		LastName:  *lastName,
		Birthday:  *birthday,
	}

	output, err := c.userUsecase.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbUser.RegisterResponse)
	res.User = UserTranslator{}.ToProto(output.User)
	return
}

func (c *userController) Get(ctx context.Context, req *pbUser.GetRequest) (res *pbUser.GetResponse, err error) {
	id, err := UserTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := user.GetUsecaseInput{
		Id: *id,
	}

	output, err := c.userUsecase.Get(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbUser.GetResponse)
	res.User = UserTranslator{}.ToProto(output.User)
	return
}

func (c *userController) Update(ctx context.Context, req *pbUser.UpdateRequest) (res *pbUser.UpdateResponse, err error) {
	id, err := UserTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	firstName, err := UserTranslator{}.ToDomainName(req.FirstName)
	if err != nil {
		firstName, _ = UserTranslator{}.ToDomainName("")
	}

	lastName, err := UserTranslator{}.ToDomainName(req.LastName)
	if err != nil {
		lastName, _ = UserTranslator{}.ToDomainName("")
	}

	birthday := &user.Birthday{}
	if req.Birthday != nil {
		birthday, _ = UserTranslator{}.ToDomainBirthday(*req.Birthday)
	}

	input := user.UpdateUsecaseInput{
		Id:        *id,
		FirstName: *firstName,
		LastName:  *lastName,
		Birthday:  *birthday,
	}

	output, err := c.userUsecase.Update(ctx, input)

	res = new(pbUser.UpdateResponse)
	res.User = UserTranslator{}.ToProto(output.User)
	return
}

func (c *userController) Terminate(ctx context.Context, req *pbUser.TerminateRequest) (res *pbUser.TerminateResponse, err error) {
	id, err := UserTranslator{}.ToDomainId(req.Id)
	if err != nil {
		return nil, err
	}

	input := user.TerminateUsecaseInput{
		Id: *id,
	}

	output, err := c.userUsecase.Terminate(ctx, input)
	if err != nil {
		return nil, err
	}

	res = new(pbUser.TerminateResponse)
	res.User = UserTranslator{}.ToProto(output.User)
	return
}
