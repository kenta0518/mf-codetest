package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
	"github.com/kenta0518/mf-codetest/pkg/domain/repository"
	"github.com/kenta0518/mf-codetest/pkg/usecase/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type User interface {
	CreateUser(ctx context.Context) (*model.User, error)
}

type userUsecase struct {
	UserRepository repository.User
	localizer      *i18n.Localizer
	tx             repository.DbTransaction
}

func NewUserUsecase(ua repository.User, localizer *i18n.Localizer, tx repository.DbTransaction) User {
	return &userUsecase{
		localizer:      localizer,
		UserRepository: ua,
		tx:             tx,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context) (*model.User, error) {
	value, err := u.tx.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		guid := uuid.NewString()
		var user = entity.User{Name: "dammy" + guid}

		newUser, err := u.UserRepository.Create(ctx, user.Name)
		if err != nil {
			c := &i18n.LocalizeConfig{MessageID: model.E0102}
			return nil, model.NewErrUnprocessable(model.E0102, u.localizer.MustLocalize(c))
		}

		return newUser, nil
	})

	if err != nil {
		return nil, err
	}

	return model.NewUser(value.(*entity.User)), nil
}
