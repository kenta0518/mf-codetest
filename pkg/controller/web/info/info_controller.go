package info

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenta0518/mf-codetest/config"
)

type InfoController struct {
	cfg *config.Config
}

func NewInfoController(cfg *config.Config) *InfoController {
	return &InfoController{
		cfg: cfg,
	}
}

func (ctl *InfoController) GetCredit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "info/credit", gin.H{})
}

func (ctl *InfoController) GetPrivacy(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "info/privacy", gin.H{})
}

func (ctl *InfoController) GetTerms(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "info/terms", gin.H{})
}
