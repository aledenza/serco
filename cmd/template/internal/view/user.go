package view

import (
	"context"
	"service/internal/service"
	"service/internal/view/schema"

	"github.com/danielgtaylor/huma/v2"
)

type UserView struct {
	service service.UserServiceI
}

func NewUserView(service service.UserServiceI) *UserView {
	return &UserView{service: service}
}

func (ua *UserView) GetUser(ctx context.Context, input *schema.UserRequest) (*schema.UserResponse, error) {
	user, success := ua.service.GetUser(ctx, input.UserId)
	if !success {
		return nil, huma.Error500InternalServerError("error getting user")
	}
	return &schema.UserResponse{
		Body: schema.User{UserId: user.UserId, FirstName: user.FirstName, SecondName: user.SecondName},
	}, nil
}

func (ua *UserView) PostUser(ctx context.Context, input *schema.UserResponse) (*schema.UserRequest, error) {
	return &schema.UserRequest{UserId: 1}, nil
}
