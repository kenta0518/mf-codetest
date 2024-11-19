package entity

import (
	"gorm.io/datatypes"
)

type GiftStockLog struct {
	LogBase
	MirrativID           string `gorm:"index:idx_mirrativ_id_stock,priority:1"`
	GiftTransactionLogID uint
	GiftTransactionLog   GiftTransactionLog
	GiftItemID           uint
	GiftItem             GiftItem
	Stock                uint `gorm:"index:idx_mirrativ_id_stock,priority:2,default:0"`
	Result               datatypes.JSONType[GiftStockResultLog]
	IsShowPopup          bool `gorm:"default:false"`
}

type GiftStockResultLog struct {
	SenderResults   []RewardContent `json:"sender_results"`
	StreamerResults []RewardContent `json:"streamer_results"`
}

func NewGiftStockLog(txLog *GiftTransactionLog, giftItem *GiftItem) *GiftStockLog {
	return &GiftStockLog{
		MirrativID:           txLog.StreamerID,
		GiftTransactionLogID: txLog.ID,
		GiftTransactionLog:   *txLog,
		GiftItemID:           giftItem.ID,
		GiftItem:             *giftItem,
		Stock:                1,
		IsShowPopup:          false,
	}
}

func NewGachaGiftStockLog(txLog *GiftTransactionLog, giftItem *GiftItem, streamerResult, senderResult []RewardContent) *GiftStockLog {
	result := GiftStockResultLog{SenderResults: senderResult, StreamerResults: streamerResult}

	return &GiftStockLog{
		MirrativID:           txLog.StreamerID,
		GiftTransactionLogID: txLog.ID,
		GiftTransactionLog:   *txLog,
		GiftItemID:           giftItem.ID,
		GiftItem:             *giftItem,
		Stock:                1,
		IsShowPopup:          false,
		Result:               datatypes.NewJSONType(result),
	}
}

func (l *GiftStockLog) UpdateStock(stock uint) {
	l.Stock = stock
	l.IsShowPopup = true
}
