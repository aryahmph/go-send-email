package controller

import (
	"go-send-email/exception"
	"go-send-email/helper"
	"go-send-email/model/web"
	"go-send-email/service"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	userCreateRequest := web.CreateUserRequest{}
	err := helper.ReadFromRequestBody(request, &userCreateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	userResponse, err := controller.UserService.Create(request.Context(), userCreateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
