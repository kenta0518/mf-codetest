package entity

type VipEffectType string

const (
	VipEffectTypeEventBattleChallenge       VipEffectType = "EventBattleChallenge"
	VipEffectTypeEventStaminaRecoveryAmount VipEffectType = "EventStaminaRecoveryAmount"
	VipEffectTypeSubQuestSkip               VipEffectType = "SubQuestSkip"
	VipEffectTypeSubQuestPlayCount          VipEffectType = "SubQuestPlayCount"
	VipEffectTypeVipLoginBonus              VipEffectType = "VipLoginBonus"
	VipEffectTypeVipDailyMission            VipEffectType = "VipDailyMission"
)

type VipEffect struct {
	SeedBase    `yaml:",inline"`
	VipID       uint          `yaml:"vipId"`
	VipEffect   VipEffectType `yaml:"vipEffect"`
	Variable    int           `yaml:"var"`
	DisplayText string        `yaml:"displayText" gorm:"size:512"`
}
