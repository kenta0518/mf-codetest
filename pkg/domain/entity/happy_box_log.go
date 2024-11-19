package entity

import (
	"gorm.io/datatypes"
)

type HappyBoxLog struct {
	LogBase
	UserID          uint
	ContentType     ContentType
	ContentID       uint
	ContentQuantity int
	IsSender        bool
	TransactionID   string
	Result          datatypes.JSONType[HappyBoxResultLog]
}

type HappyBoxResultLog struct {
	Results []RewardContent `json:"results"`
}

func NewSenderHappyBoxLog(userId uint, content RewardContent, result []RewardContent, transactionID string) *HappyBoxLog {
	return &HappyBoxLog{
		UserID:          userId,
		IsSender:        true,
		ContentType:     content.ContentType,
		ContentID:       content.ContentID,
		ContentQuantity: content.ContentQuantity,
		TransactionID:   transactionID,
		Result:          datatypes.NewJSONType(HappyBoxResultLog{Results: result}),
	}
}

func NewStreamerHappyBoxLog(userId uint, content RewardContent, result []RewardContent, transactionID string) *HappyBoxLog {
	return &HappyBoxLog{
		UserID:          userId,
		IsSender:        false,
		ContentType:     content.ContentType,
		ContentID:       content.ContentID,
		ContentQuantity: content.ContentQuantity,
		TransactionID:   transactionID,
		Result:          datatypes.NewJSONType(HappyBoxResultLog{Results: result}),
	}
}
