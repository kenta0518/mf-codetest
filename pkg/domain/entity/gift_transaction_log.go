package entity

import (
	"time"
)

type GiftTransactionState int

const (
	GIFT_TRANSACTION_STATE_OPEN = iota
	GIFT_TRANSACTION_STATE_CLOSE
	GIFT_TRANSACTION_STATE_CANCEL
)

type GiftTransactionLog struct {
	LogBase
	ItemID           uint
	ItemNum          uint
	StreamerID       string
	StreamerName     string
	StreamerImageUrl string
	SenderID         string `gorm:"index:idx_sender_id_state_created_at_DESC,priority:1"`
	SenderName       string
	SenderImageUrl   string
	LiveID           string
	TransactionID    string `gorm:"index"`
	OpenedAt         string
	State            GiftTransactionState `gorm:"index:idx_sender_id_state_created_at_DESC,priority:2"`
	CreatedAt        time.Time            `gorm:"index:idx_sender_id_state_created_at_DESC,priority:3,sort:desc"`
}

func NewGiftTransactionLog(itemID uint, itemNum uint, streamerID string, streamerName string, streamerImageUrl string,
	senderID string, senderName string, senderImageUrl string, liveID string, transactionID string, openedAt string) *GiftTransactionLog {
	return &GiftTransactionLog{
		ItemID:           itemID,
		ItemNum:          itemNum,
		StreamerID:       streamerID,
		StreamerName:     streamerName,
		StreamerImageUrl: streamerImageUrl,
		SenderID:         senderID,
		SenderName:       senderName,
		SenderImageUrl:   senderImageUrl,
		LiveID:           liveID,
		TransactionID:    transactionID,
		OpenedAt:         openedAt,
		State:            GIFT_TRANSACTION_STATE_OPEN,
	}
}

func (log *GiftTransactionLog) Close() {
	log.State = GIFT_TRANSACTION_STATE_CLOSE
}

func (log *GiftTransactionLog) Cancel() {
	log.State = GIFT_TRANSACTION_STATE_CANCEL
}

func (log *GiftTransactionLog) IsCanceled() bool {
	return log.State == GIFT_TRANSACTION_STATE_CANCEL
}

func (log *GiftTransactionLog) IsClosed() bool {
	return log.State == GIFT_TRANSACTION_STATE_CLOSE
}
