package repository

import (
	"context"
	"go-send-email/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) domain.User
	FindByEmail(ctx context.Context, userEmail string) (domain.User, error)
}
