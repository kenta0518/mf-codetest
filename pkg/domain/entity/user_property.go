package entity

import (
	"time"

	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

type TutorialFlag uint

const (
	InitGold     = 1000
	InitStarRank = 1
)

const (
	// iota を使ってひと桁ずつシフトする
	A uint = 1 << iota // 0000001
	B                  // 0000010
	C                  // 0000100
	D                  // 0001000
	E                  // 0010000
	F                  // 0100000
	G                  // 1000000
)

const (
	INITIAL_STAMINA       = 15
	INITIAL_EVENT_STAMINA = 10
)

type UserProperty struct {
	gorm.Model
	UserID             uint
	Gold               uint
	StarRank           uint
	RankExp            uint
	BeforeStarRank     uint
	BestScore          uint
	TutorialConfirmed  TutorialFlag
	LastStaminaResetAt *DateTime
	TutorialFinished   *DateTime
	// 一旦、配信/非配信は考えない
	EventStamina uint
}

// フラグをオンにする
func (u *UserProperty) SetTutorialConfirmed(flag TutorialFlag) {
	u.TutorialConfirmed |= flag
}

// フラグをオフ(反転)にする
func (u *UserProperty) PutTutorialConfirmed(flag TutorialFlag) {
	u.TutorialConfirmed ^= flag
}

func (u *UserProperty) UpdateBestScore(score uint) bool {
	if score > u.BestScore {
		u.BestScore = score
		return true
	}
	return false
}

// スタミナリセット
func (u *UserProperty) ResetStamina(logintime DateTime, eventStaminaBuff uint) bool {
	//初回ログイン
	if u.LastStaminaResetAt == nil {
		u.EventStamina = INITIAL_EVENT_STAMINA + eventStaminaBuff
		u.LastStaminaResetAt = &logintime
		return true
	}

	//0時のリセット
	am0 := now.With(logintime.Time).BeginningOfDay()
	// LastStaminaResetAt < am0 && am0 < logintime
	if u.LastStaminaResetAt.Time.Before(am0) && am0.Before(logintime.Time) {
		u.EventStamina = INITIAL_EVENT_STAMINA + eventStaminaBuff
		u.LastStaminaResetAt = &logintime
		return true
	}

	//12時のリセット
	pm0 := am0.Add(12 * time.Hour)
	// LastStaminaResetAt < pm0 && pm0 < logintime
	if u.LastStaminaResetAt.Time.Before(pm0) && pm0.Before(logintime.Time) {
		u.EventStamina = INITIAL_EVENT_STAMINA + eventStaminaBuff
		u.LastStaminaResetAt = &logintime
		return true
	}

	return false
}
