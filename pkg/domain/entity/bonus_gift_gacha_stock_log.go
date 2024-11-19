package entity

import (
	"gorm.io/datatypes"
)

type BonusGiftGachaStockLog struct {
	LogBase
	MirrativID           string `gorm:"index:idx_mirrativ_id_stock,priority:1"`
	GiftTransactionLogID uint
	GiftTransactionLog   GiftTransactionLog
	GiftItemID           uint
	GiftItem             GiftItem
	Stock                uint `gorm:"index:idx_mirrativ_id_stock,priority:2"`
	ProfileBannerID      uint
	ProfileDegree1ID     uint
	ProfileDegree2ID     uint
	ProfileDegree3ID     uint
	Result               datatypes.JSONType[BonusGiftGachaResultLog]
}

type BonusGiftGachaResultLog struct {
	Results []RewardContent `json:"results"`
}

func NewBonusGiftGachaStockLog(txLog *GiftTransactionLog, giftItem *GiftItem, bannerID, degree1ID, degree2ID, degree3ID uint, result []RewardContent) *BonusGiftGachaStockLog {
	resultLog := BonusGiftGachaResultLog{Results: result}

	return &BonusGiftGachaStockLog{
		MirrativID:           txLog.SenderID,
		GiftTransactionLogID: txLog.ID,
		GiftTransactionLog:   *txLog,
		GiftItemID:           giftItem.ID,
		GiftItem:             *giftItem,
		ProfileBannerID:      bannerID,
		ProfileDegree1ID:     degree1ID,
		ProfileDegree2ID:     degree2ID,
		ProfileDegree3ID:     degree3ID,
		Stock:                1,
		Result:               datatypes.NewJSONType(resultLog),
	}
}
