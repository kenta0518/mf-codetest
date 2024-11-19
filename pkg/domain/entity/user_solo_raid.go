package entity

import "time"

const (
	SoloRaidAttackStamina       uint = 1
	SoloRaidOverAttackStamina   uint = 3
	SoloRaidStaminaMax          uint = 10
	SoloRaidStaminaRecoveryTime      = 10 * time.Minute
)

type UserSoloRaid struct {
	UserResourceBase
	SoloRaid     SoloRaid  `gorm:"foreignKey:ResourceID"`
	BossLevel    uint      `gorm:"default:1"`
	Damage       uint      `gorm:"default:0"`
	KillCount    uint      `gorm:"default:0"`
	AllClearedAt *DateTime `gorm:"default:null"`
	Stamina      uint      `gorm:"default:0"`
	CheckedAt    DateTime
}

func NewUserSoloRaid(userID, soloRaidID uint, check time.Time) *UserSoloRaid {
	return &UserSoloRaid{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: soloRaidID,
		},
		BossLevel:    1,
		Damage:       0,
		KillCount:    0,
		AllClearedAt: nil,
		Stamina:      SoloRaidStaminaMax,
		CheckedAt:    DateTime{Time: check},
	}
}

func (u UserSoloRaid) IsAllCleared() bool {
	return u.AllClearedAt != nil
}

func (u *UserSoloRaid) AllCreal(now time.Time) {
	u.AllClearedAt = &DateTime{Time: now}
}

func (u *UserSoloRaid) AddDamage(damage uint) {
	u.Damage += damage
}

func (u *UserSoloRaid) LevelUp() {
	u.BossLevel++
	u.KillCount++
	u.Damage = 0
}

func (u *UserSoloRaid) RecoveryStamina(now time.Time) bool {
	if u.Stamina == SoloRaidStaminaMax {
		return false
	}

	// 前回のチェックからの経過時間を元にStaminateを10分ごとに1回復
	elapsed := now.Sub(u.CheckedAt.Time)
	recovery := uint(elapsed / SoloRaidStaminaRecoveryTime)
	if recovery == 0 {
		return false
	}

	// スタミナを回復
	u.Stamina += recovery
	// 回復数*回復時間をCheckedAtに加算
	checkTime := u.CheckedAt.Add(SoloRaidStaminaRecoveryTime * time.Duration(recovery))

	// スタミナが最大値を超えた場合は最大値にする
	// 最大値の時はチェック時間を更新
	if u.Stamina > SoloRaidStaminaMax {
		u.Stamina = SoloRaidStaminaMax
		checkTime = now
	}

	u.CheckedAt = DateTime{Time: checkTime}

	return true
}

func (u *UserSoloRaid) CheckStamina(overAttack bool) bool {
	if overAttack {
		return u.Stamina >= SoloRaidOverAttackStamina
	}
	return u.Stamina >= SoloRaidAttackStamina
}

func (u *UserSoloRaid) UseStamina(overAttack bool, now time.Time) bool {
	sub := SoloRaidAttackStamina
	if overAttack {
		sub = SoloRaidOverAttackStamina
	}

	// スタミナが足りない
	if u.Stamina < sub {
		return false
	}

	// スタミナを消費
	// スタミナが最大値の場合はチェック時間を更新
	if u.Stamina == SoloRaidStaminaMax {
		u.CheckedAt = DateTime{Time: now}
	}
	u.Stamina -= sub

	return true
}
