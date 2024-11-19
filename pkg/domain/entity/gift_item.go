package entity

type GiftType string

/*
Gacha	// ガチャ
InGameAdvantagePieceChange	// 有利属性ピース生成
InGameHealPieceChange	// 回復ピース生成
InGameEncoreTime	// アンコールタイム
InGameRandomGift	// チャンスギフト
InGameMeteor	// メテオ
*/

const (
	GiftTypeGacha                      GiftType = "Gacha"                      // ガチャ
	GiftTypeInGameAdvantagePieceChange GiftType = "InGameAdvantagePieceChange" // 有利属性ピース生成
	GiftTypeInGameHealPieceChange      GiftType = "InGameHealPieceChange"      // 回復ピース生成
	GiftTypeInGameEncoreTime           GiftType = "InGameEncoreTime"           // アンコールタイム
	GiftTypeInGameRandomGift           GiftType = "InGameRandomGift"           // チャンスギフト
	GiftTypeInGameMeteor               GiftType = "InGameMeteor"               // メテオ
)

type GiftItem struct {
	SeedBase     `yaml:",inline"`
	ItemID       uint     `yaml:"itemId"`
	GiftType     GiftType `yaml:"type"`
	Name         string   `yaml:"name"`
	EffectKind   int      `yaml:"effectKind"`
	EffectValue1 uint     `yaml:"effectValue1"`
	EffectValue2 uint     `yaml:"effectValue2"`
	Coin         int      `yaml:"coin"`
}
