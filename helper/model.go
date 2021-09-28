package helper

import (
	"go-send-email/model/domain"
	"go-send-email/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Email: user.Email,
	}
}
