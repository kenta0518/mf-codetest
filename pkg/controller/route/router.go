package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	_ "github.com/kenta0518/mf-codetest/docs"
)

func Route(
	route *gin.Engine,
	cfg *config.Config,
	localizer *i18n.Localizer,
) {

	// ヘルスチェック
	route.GET("/", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"status": "ok"}) })

	// 認証なし
	//route.POST("/api/sessions", cs.CreateSession)
	//route.GET("/api/resourceversion", resource.GetResourceVersion)

}
