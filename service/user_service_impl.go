package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go-send-email/exception"
	"go-send-email/helper"
	"go-send-email/model/domain"
	"go-send-email/model/web"
	"go-send-email/repository"
	"gorm.io/gorm"
	"log"
	"strings"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	MailService    MailService
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, mailService MailService, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository, MailService: mailService, Validate: validate}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.CreateUserRequest) (web.UserResponse, error) {
	// Validate
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	// Convert to struct
	user := domain.User{
		Email: strings.Trim(request.Email, " "),
	}

	// Check email
	_, err = service.UserRepository.FindByEmail(ctx, user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("err")
		return web.UserResponse{}, exception.ErrAlreadyExist
	}

	// Save
	user = service.UserRepository.Save(ctx, user)

	// Send Mail
	sendMail := domain.SendMail{
		Subject: "Welcome to AryaHmph Company !",
		Message: "Lorem ipsum sit dolor amet.",
	}
	sendMail.To = append(sendMail.To, user.Email)
	go service.MailService.Send(sendMail)

	return helper.ToUserResponse(user), nil
}
