package repository

import (
	"context"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
	"github.com/kenta0518/mf-codetest/pkg/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &userRepository{
		db: db,
	}
}

func (t *userRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	tx, ok := GetTx(ctx)

	if !ok {
		return nil, repository.ErrTx
	}

	user := entity.User{Name: name}
	err := tx.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}
