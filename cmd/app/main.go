package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
	"github.com/kenta0518/mf-codetest/pkg/controller"
	"github.com/kenta0518/mf-codetest/pkg/controller/route"
	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
	"github.com/kenta0518/mf-codetest/pkg/infra"
	"github.com/kenta0518/mf-codetest/pkg/infra/repository"
	"github.com/kenta0518/mf-codetest/pkg/usecase"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/fx"
	"golang.org/x/text/language"
	"gorm.io/gorm"

	"github.com/gin-contrib/secure"
)

//	@title						MFテスト API
//	@version					1.0
//	@description				MFテストのテストAPI

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	handler := initHandler(cfg)

	mysql := infra.NewMySqlConnector(cfg)
	localizer := newLocalizer()

	migrate(mysql.DB)

	app := fx.New(
		fx.Supply(cfg, localizer, mysql.DB, handler),
		repository.Modules(),
		usecase.Modules(),
		controller.Modules(),
		fx.Invoke(
			lifecycle,
			route.Route,
			func(r *gin.Engine) {},
		),
	)

	app.Run()
}

func initHandler(cfg *config.Config) *gin.Engine {

	securityCfg := secure.DefaultConfig()
	securityCfg.CustomFrameOptionsValue = "SAMEORIGIN"
	securityCfg.SSLRedirect = false
	securityCfg.ContentSecurityPolicy = ""

	handler := gin.New()

	handler.Use(secure.New(securityCfg))
	handler.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Robots-Tag", "noindex, nofollow, nosnippet, noarchive")
	})
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	return handler
}

func newLocalizer() *i18n.Localizer {
	bundle := i18n.NewBundle(language.Japanese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile("language/ja.toml")
	if err != nil {
		panic(err)
	}

	localizer := i18n.NewLocalizer(bundle)
	return localizer
}

func lifecycle(lc fx.Lifecycle, cfg *config.Config, handler *gin.Engine) {
	srv := &http.Server{Addr: ":8888", Handler: handler}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go (func() {
				_ = srv.ListenAndServe()
			})()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(entity.Entity()...)

	if err != nil {
		panic(err)
	}
}
