package entity

type ItemType string

const (
	ITEM_CHARACTER_ENHANCE ItemType = "CharaEnhance"  // キャラ強化
	ITEM_HAPPY_BOX         ItemType = "HappyBox"      // ハッピーボックス
	ITEM_SUPER_HAPPY_BOX   ItemType = "SuperHappyBox" // ハッピーボックス
	ITEM_GACHA_TICKET      ItemType = "GachaTicket"   // ガチャチケット
)

type Item struct {
	SeedBase `yaml:",inline"`
	Name     string   `yaml:"name"`
	ImageID  string   `yaml:"imageId"`
	ItemType ItemType `yaml:"itemType"`
	Value1   int      `yaml:"value1"`
}
