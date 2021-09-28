package web

type CreateUserRequest struct {
	Email string `validate:"required,email" json:"email"`
}

type SendMailRequest struct {
	Message string `validate:"required" json:"message"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}