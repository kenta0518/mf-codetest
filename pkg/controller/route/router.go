package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
	_ "github.com/kenta0518/mf-codetest/docs"
	"github.com/kenta0518/mf-codetest/pkg/controller"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route(
	route *gin.Engine,
	localizer *i18n.Localizer,
	cfg *config.Config,
	user *controller.UserController,
	transaction *controller.TransactionController,
) {
	// ヘルスチェック
	route.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "ok"}) })
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.POST("/api/users", user.CreateUser)
	route.POST("/transactions", transaction.CreateTransaction)
	//route.POST("/api/sessions", cs.CreateSession)
	//route.GET("/api/resourceversion", resource.GetResourceVersion)

}
