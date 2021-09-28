package repository

import (
	"context"
	"go-send-email/helper"
	"go-send-email/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.User) domain.User {
	err := repository.DB.WithContext(ctx).Create(&user).Error
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, userEmail string) (domain.User, error) {
	var user domain.User
	err := repository.DB.WithContext(ctx).Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
