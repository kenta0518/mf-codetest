package entity

func Entity() []any {
	return concatSlices([]any{
		// User
		&User{},
		&UserDeck{},
		&UserProperty{},
		&UserPresent{},
		&UserOshirase{},
		&UserMission{},
		&UserLoginState{},
		&UserLiveState{},
		&UserInterstitialBanner{},
		&UserVipProperty{},
		&UserGiftProperty{},
		&UserVip{},
		&UserLoginBonusReward{},
		&UserExchangeLineup{},

		// Ranking
		&UserRanking{},

		// その他
		&Oshirase{},
		&Present{},
		&Admin{},
		&InterstitialBanner{},
		&NewsHeadline{},
		&RectangleBanner{},

		// ログ
		&GiftTransactionLog{},
		&GiftStockLog{},
		&GachaLog{},
		&HappyBoxLog{},
		&GiftGachaLog{},
		&BonusGiftGachaStockLog{},
		&ExchangeLog{},
	},
		// ユーザリソース
		UserResource(),

		// シード
		Seed(),
	)
}

func UserResource() []any {
	return []any{
		&UserBonus{},
		&UserCharacterCard{},
		&UserItem{},
		&UserPrizeSticker{},

		// MainQuest
		&UserMainQuestGroup{},
		&UserMainQuest{},

		// SubQuest
		&UserSubQuestGroup{},
		&UserSubQuest{},

		// MainAdv
		&UserMainAdv{},

		// PanelMission
		&UserPanelMission{},

		// SoloRaid
		&UserSoloRaid{},

		// EventSeason
		&UserEventSeason{},
		&UserEventQuest{},
	}
}

func Seed() []any {
	return []any{
		&GiftItem{},
		&HappyBox{},
		&ResourceVersion{},
		&Gacha{},
		&GachaPrize{},
		&Item{},
		&Bonus{},
		&StarRankTable{},
		&MainAdv{},
		&MainAdvGroup{},
		&VoltageRewardTable{},
		&Vip{},
		&VipEffect{},
		&VipGift{},
		&Exchange{},
		&ExchangeLineup{},
		&ExchangeContent{},
		&PrizeSticker{},

		// Stars2
		// Character
		&CharacterCard{},
		&CharacterEnhance{},
		&CharacterItemConvert{},

		// Quest
		&QuestReward{},
		&QuestScoreRank{},
		// MainQuest
		&MainQuestGroup{},
		&MainQuest{},
		// SubQuest
		&SubQuestGroup{},
		&SubQuest{},

		// LoginBonus
		&LoginBonus{},
		&LoginBonusReward{},
		&LoginBonusRewardGroup{},

		// Mission
		&MissionGroup{},
		&Mission{},
		&PanelMissionGroup{},
		&PanelMission{},
		&MissionReward{},

		// SoloRaid
		&SoloRaid{},
		&SoloRaidBoss{},
		&SoloRaidReward{},

		// EventSeason
		&EventSeason{},
		&RankingGroup{},
		&RankingRange{},
		&RankingReward{},
		&EventQuest{},
	}
}

// Seed Union
type SeedType interface {
	SeedModule()
}

// Log Union
type LogType interface {
	LogModule()
	GetID() uint
	IsEmpty() bool
}

type UserResourceType interface {
	UserResourceModule()
	GetID() uint
	IsEmpty() bool
}

func concatSlices(slices ...[]any) []any {
	var result []any
	for _, slice := range slices {
		result = append(result, slice...)
	}
	return result
}
