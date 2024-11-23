package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
	"github.com/kenta0518/mf-codetest/pkg/usecase"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UserController struct {
	controllerBase
	userUsecase usecase.User
}

func NewUserController(ur usecase.User, cfg *config.Config, lc *i18n.Localizer) *UserController {
	return &UserController{
		userUsecase:    ur,
		controllerBase: controllerBase{cfg: cfg, localizer: lc},
	}
}

// Result godoc
//
// @Summary ユーザー作成
// @Description ユーザーを作成します
// @Accept json
// @Produce json
// @Success 200 {object} model.User
// @Failure 400 {object} model.AppError
// @Router /api/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	user, err := c.userUsecase.CreateUser(ctx)
	if err != nil {
		// エラーが発生した場合はステータスコードとエラーメッセージを返す
		apperr := c.toAppError(err)
		ctx.JSON(apperr.StatusCode, gin.H{"error": apperr})
		return
	}

	// 正常時のレスポンスを返す
	ctx.JSON(http.StatusOK, user)
}
