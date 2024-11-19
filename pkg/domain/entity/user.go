package entity

import (
	"errors"
	"time"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

var (
	ErrLackOfResources       = errors.New("lack of resources")
	ErrUserEquipmentNotFound = errors.New("user equipment not found")
)

const (
	Tutee     = iota // ゲストユーザー(=チュートリアル終了まで)
	Player           // 通常ユーザ
	SuperUser        // スーパーユーザ
	Banned           // バン
	Deleted          // 削除
)

type User struct {
	gorm.Model
	Name                string `gorm:"index;size:256"`
	MirrativID          string `gorm:"uniqueIndex;size:256"`
	ProfileUrl          string
	UserKind            uint
	TimeDifference      time.Duration
	Property            UserProperty
	LoginState          UserLoginState
	LiveState           UserLiveState
	Items               []UserItem
	Bonuses             []UserBonus
	Oshirases           []UserOshirase
	Missions            []UserMission
	InterstitialBanners []UserInterstitialBanner
	Vips                []UserVip
	LoginBonusRewards   []UserLoginBonusReward
	ExchangeLineups     []UserExchangeLineup
	PrizeStickers       []UserPrizeSticker

	// TODO: STARS2
	// 処理の見直しをしたいのもあるので、STARS2のResoruceを以下に追加していく
	MainAdvs        []UserMainAdv
	CharacterCards  []UserCharacterCard
	Decks           []UserDeck
	MainQuestGroups []UserMainQuestGroup
	MainQuests      []UserMainQuest
	SubQuestGroups  []UserSubQuestGroup
	SubQuests       []UserSubQuest
	PanelMissions   []UserPanelMission
}

func (u User) IsEmpty() bool {
	return u.ID == 0
}

func (u User) IsSuperUser() bool {
	return u.UserKind == SuperUser
}

func (u *User) AddStarRankExp(value uint) {
	u.Property.RankExp += value
}

func (u *User) UpdateStarRank(value uint) {
	u.Property.StarRank = value
}

func (u *User) GainGold(value uint) bool {
	u.Property.Gold += value

	return true
}

func (u *User) ConsumeGold(value uint) bool {
	if u.Property.Gold < value {
		return false
	}
	u.Property.Gold -= value

	return true
}

func (u *User) GainItem(itemID uint, quantity int) *UserItem {
	idx := funk.IndexOf(u.Items, func(item UserItem) bool { return item.ResourceID == itemID })
	if idx < 0 {
		userItem := NewUserItem(u.ID, itemID, quantity)
		u.Items = append(u.Items, *userItem)

		return &u.Items[len(u.Items)-1]
	}

	u.Items[idx].Gain(quantity)

	return &u.Items[idx]
}

func (u *User) GainPrizeSticker(prizeStickerID uint, quantity int) *UserPrizeSticker {
	idx := funk.IndexOf(u.PrizeStickers, func(up UserPrizeSticker) bool { return up.ResourceID == prizeStickerID })
	if idx < 0 {
		userPrizeSticker := NewUserPrizeSticker(u.ID, prizeStickerID, quantity)
		u.PrizeStickers = append(u.PrizeStickers, *userPrizeSticker)

		return &u.PrizeStickers[len(u.PrizeStickers)-1]
	}

	u.PrizeStickers[idx].Gain(quantity)

	return &u.PrizeStickers[idx]
}

func (u *User) GainBonus(bonusID uint, quantity int) *UserBonus {
	idx := funk.IndexOf(u.Bonuses, func(bonus UserBonus) bool { return bonus.ResourceID == bonusID })
	if idx < 0 {
		userBonus := NewUserBonus(u.ID, bonusID, quantity)
		u.Bonuses = append(u.Bonuses, *userBonus)

		return &u.Bonuses[len(u.Bonuses)-1]
	}

	u.Bonuses[idx].Gain(quantity)

	return &u.Bonuses[idx]
}

func (u *User) GainCharacterCard(characterCardID uint) *UserCharacterCard {
	idx := funk.IndexOf(u.CharacterCards, func(card UserCharacterCard) bool { return card.ResourceID == characterCardID })
	if idx < 0 {
		userCharacterCard := NewUserCharacterCard(u.ID, characterCardID)
		u.CharacterCards = append(u.CharacterCards, *userCharacterCard)

		return &u.CharacterCards[len(u.CharacterCards)-1]
	}

	u.CharacterCards[idx].UpAwake()

	return &u.CharacterCards[idx]
}

func (u *User) ConsumeItem(itemID uint, quantity int) (*UserItem, error) {
	idx := funk.IndexOf(u.Items, func(item UserItem) bool { return item.ResourceID == itemID })
	if idx < 0 {
		return nil, ErrLackOfResources
	}

	userItem := &u.Items[idx]
	if ok := userItem.Consume(quantity); !ok {
		return nil, ErrLackOfResources
	}

	return userItem, nil
}

func (u *User) GetItem(itemID uint) *UserItem {
	idx := funk.IndexOf(u.Items, func(item UserItem) bool { return item.ResourceID == itemID })

	if idx < 0 {
		return nil
	}

	return &u.Items[idx]
}

func (u *User) GetBonus(bonusID uint) *UserBonus {
	idx := funk.IndexOf(u.Bonuses, func(bonus UserBonus) bool { return bonus.ResourceID == bonusID })

	if idx < 0 {
		return nil
	}

	return &u.Bonuses[idx]
}

func (u *User) GetItems(effectType ItemType) []UserItem {
	items := []UserItem{}
	for _, item := range u.Items {
		if item.Item.ItemType == effectType {
			items = append(items, item)
		}
	}

	return items
}

func (u *User) GetAdv(advId uint) *UserMainAdv {
	idx := funk.IndexOf(u.MainAdvs, func(adv UserMainAdv) bool { return adv.ResourceID == advId })
	if idx < 0 {
		return nil
	}
	return &u.MainAdvs[idx]
}

func (u *User) GetOshirase(oshiraseID uint) *UserOshirase {
	idx := funk.IndexOf(u.Oshirases, func(oshirase UserOshirase) bool { return oshirase.OshiraseID == oshiraseID })
	if idx < 0 {
		return nil
	}
	return &u.Oshirases[idx]
}

func (u *User) ChangeNormalUser() {
	u.UserKind = Player
}

func (u *User) UpdateUserKind(kind uint) {
	u.UserKind = kind
}

func (u *User) GetScore(value uint) bool {
	u.Property.BestScore = value

	return true
}

func (u User) GetUserVipByExp(exp uint) *UserVip {
	if len(u.Vips) == 0 {
		return nil
	}

	// goalが昇順になってる前提
	var userVip *UserVip
	for idx := range u.Vips {
		if u.Vips[idx].Vip.Goal <= exp {
			userVip = &u.Vips[idx]
		} else {
			break
		}
	}

	return userVip
}

func (u User) GetLoginBonusRewards(loginBonusID uint) []UserLoginBonusReward {
	value := funk.Filter(u.LoginBonusRewards, func(ur UserLoginBonusReward) bool { return ur.LoginBonusReward.LoginBonusID == loginBonusID })
	if value == nil {
		return nil
	}
	return value.([]UserLoginBonusReward)
}
