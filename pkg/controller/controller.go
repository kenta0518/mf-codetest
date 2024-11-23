package controller

import (
	"github.com/kenta0518/mf-codetest/config"
	"github.com/kenta0518/mf-codetest/pkg/usecase/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type controllerBase struct {
	cfg       *config.Config
	localizer *i18n.Localizer
}

func (c *controllerBase) toAppError(err error) *model.AppError {
	switch apperr := err.(type) {
	case *model.AppError:
		return apperr
	default:
		cf := &i18n.LocalizeConfig{MessageID: model.E9999}
		return model.NewErrInternalServerError(model.E9999, c.localizer.MustLocalize(cf))
	}
}
