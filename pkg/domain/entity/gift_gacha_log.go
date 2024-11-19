package entity

import (
	"gorm.io/datatypes"
)

type GiftGachaLog struct {
	LogBase
	MirrativID           string `gorm:"Index;size:256"`
	GiftItemID           uint
	GiftItem             GiftItem
	GiftTransactionLogID uint
	GiftTransactionLog   GiftTransactionLog
	IsSender             bool
	Result               datatypes.JSONType[GiftGachaResultLog]
}

type GiftGachaResultLog struct {
	Results []RewardContent `json:"results"`
}

func NewSenderGiftGachaLog(mirrativID string, giftItem GiftItem, txLog GiftTransactionLog, results []RewardContent) *GiftGachaLog {
	return &GiftGachaLog{
		MirrativID:           mirrativID,
		GiftItemID:           giftItem.ID,
		GiftItem:             giftItem,
		GiftTransactionLogID: txLog.ID,
		GiftTransactionLog:   txLog,
		IsSender:             true,
		Result:               datatypes.NewJSONType(GiftGachaResultLog{Results: results}),
	}
}

func NewStreamerGiftGachaLog(mirrativID string, giftItem GiftItem, txLog GiftTransactionLog, results []RewardContent) *GiftGachaLog {
	return &GiftGachaLog{
		MirrativID:           mirrativID,
		GiftItemID:           giftItem.ID,
		GiftItem:             giftItem,
		GiftTransactionLogID: txLog.ID,
		GiftTransactionLog:   txLog,
		IsSender:             false,
		Result:               datatypes.NewJSONType(GiftGachaResultLog{Results: results}),
	}
}
