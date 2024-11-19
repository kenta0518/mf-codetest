package controller

import (
	"errors"
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/kenta0518/mf-codetest/pkg/usecase/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func createLocalizer() (*i18n.Localizer, error) {
	bundle := i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile("../../language/ja.toml")
	if err != nil {
		return nil, err
	}

	localizer := i18n.NewLocalizer(bundle)
	return localizer, nil
}

func Test_controllerBase_toAppError(t *testing.T) {
	lc, err := createLocalizer()
	if err != nil {
		t.Fatalf("i18n file error: '%s' ", err)
		return
	}

	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want *model.AppError
	}{
		{
			name: "エラー生成その1",
			args: args{
				err: model.NewAppError(500, "code", "msg"),
			},
			want: model.NewAppError(500, "code", "msg"),
		},
		{
			name: "エラー生成その2",
			args: args{
				err: errors.New("err"),
			},
			want: model.NewErrInternalServerError(model.E9999, lc.MustLocalize(&i18n.LocalizeConfig{MessageID: model.E9999})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controllerBase{
				localizer: lc,
			}
			if got := c.toAppError(tt.args.err); !reflect.DeepEqual(got.ErrorMessage, tt.want.ErrorMessage) {
				t.Errorf("controllerBase.toAppError() = %v, want %v", got, tt.want)
			}
		})
	}
}
