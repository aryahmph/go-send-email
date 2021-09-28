package controller

import "net/http"

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request)
}
