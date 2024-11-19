package entity

import (
	"gorm.io/datatypes"
)

type GachaLog struct {
	LogBase
	UserID          uint
	ContentID       uint
	ContentType     ContentType
	ContentQuantity int
	Count           uint
	GachaID         uint
	Result          datatypes.JSONType[GachaResultLog]
}

type GachaResultLog struct {
	UserResults []RewardContent `json:"user_results"`
}

func NewGachaLog(userId uint, count uint, gachaId uint, content RewardContent, userResult []RewardContent) *GachaLog {
	result := GachaResultLog{UserResults: userResult}
	return &GachaLog{
		UserID:          userId,
		ContentID:       content.ContentID,
		ContentType:     content.ContentType,
		ContentQuantity: content.ContentQuantity,
		Count:           count,
		GachaID:         gachaId,
		Result:          datatypes.NewJSONType(result),
	}
}
