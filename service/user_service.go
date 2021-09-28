package service

import (
	"context"
	"go-send-email/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.CreateUserRequest) (web.UserResponse, error)
}
