package entity

type ContentType string

const (
	CONTENT_TYPE_NONE          ContentType = "None"
	CONTENT_TYPE_ITEM          ContentType = "Item"
	CONTENT_TYPE_GOLD          ContentType = "Gold"
	CONTENT_TYPE_STAREXP       ContentType = "StarExp"
	CONTENT_TYPE_BONUS         ContentType = "Bonus"
	CONTENT_TYPE_EVENT_STAMINA ContentType = "EventStamina"
	CONTENT_TYPE_PRIZE_STICKER ContentType = "PrizeSticker"

	// Stars2
	CONTENT_TYPE_CHARACTER_CARD ContentType = "CharacterCard"
)

var ContentTypeList = []ContentType{
	CONTENT_TYPE_NONE,
	CONTENT_TYPE_ITEM,
	CONTENT_TYPE_GOLD,
	CONTENT_TYPE_STAREXP,
	CONTENT_TYPE_BONUS,
	CONTENT_TYPE_EVENT_STAMINA,
	CONTENT_TYPE_PRIZE_STICKER,
	CONTENT_TYPE_CHARACTER_CARD,
}

type RewardContent struct {
	ContentType     ContentType `yaml:"contentType" json:"contentType"`
	ContentID       uint        `yaml:"contentId" json:"contentId"`
	ContentQuantity int         `yaml:"contentQuantity" json:"contentQuantity"`
}
