package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
	"github.com/kenta0518/mf-codetest/pkg/usecase"
	"github.com/kenta0518/mf-codetest/pkg/usecase/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type TransactionController struct {
	controllerBase
	transactionUsecase usecase.Transaction
}

func NewTransactionController(tu usecase.Transaction, cfg *config.Config, lc *i18n.Localizer) *TransactionController {
	return &TransactionController{
		controllerBase:     controllerBase{cfg: cfg, localizer: lc},
		transactionUsecase: tu,
	}
}

// CreateTransaction godoc
//
// @Summary	取引登録
// @Description	取引登録処理が仕様を満たしているかテストする。
// @Accept			json
// @Produce		json
// @Param			transaction	body	Transaction	true	"取引情報"
// @Success		200	{object}	model.Transaction
// @Failure		400	{object}	model.AppError
// @Router			/transactions [post]
func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req Transaction
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apperr := c.toAppError(err)
		ctx.JSON(apperr.StatusCode, gin.H{"error": apperr})
		return
	}

	// CreateTransaction メソッドに引数を渡す
	transaction, err := c.transactionUsecase.CreateTransaction(ctx, req.UserID, req.Amount, req.Description)
	if err != nil {
		cf := &i18n.LocalizeConfig{MessageID: model.E0201}
		apperr := model.NewErrPaymentRequired(model.E0201, c.localizer.MustLocalize(cf))
		ctx.JSON(apperr.StatusCode, gin.H{"error": apperr})
		return
	}

	// 正常時のレスポンスを返す
	ctx.JSON(http.StatusOK, transaction)
}

type Transaction struct {
	UserID      int    `json:"user_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
