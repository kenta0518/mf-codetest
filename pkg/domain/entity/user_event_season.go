package entity

import "time"

const (
	EventSeasonStaminaMax          uint = 10
	EventSeasonStaminaRecoveryTime      = 10 * time.Minute
)

type UserEventSeason struct {
	UserResourceBase
	EventSeason EventSeason `gorm:"foreignKey:ResourceID"`
	EventPoint  uint
	Stamina     uint
	CheckedAt   DateTime
}

func NewUserEventSeason(userID, eventSeasonID uint, now time.Time) *UserEventSeason {
	return &UserEventSeason{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: eventSeasonID,
		},
		EventPoint: 0,
		Stamina:    EventSeasonStaminaMax,
		CheckedAt:  DateTime{Time: now},
	}
}

func (u *UserEventSeason) AddEventPoint(point uint) {
	u.EventPoint += point
}

func (u *UserEventSeason) UseStamina(now time.Time) bool {
	if u.Stamina == 0 {
		return false
	}

	if u.Stamina == EventSeasonStaminaMax {
		u.CheckedAt = DateTime{Time: now}
	}

	u.Stamina--
	return true
}

func (u *UserEventSeason) RecoveryStamina(now time.Time) bool {
	if u.Stamina == EventSeasonStaminaMax {
		return false
	}

	elapsed := now.Sub(u.CheckedAt.Time)
	recovery := uint(elapsed / EventSeasonStaminaRecoveryTime)
	if recovery == 0 {
		return false
	}

	u.Stamina += recovery
	checkTime := u.CheckedAt.Add(EventSeasonStaminaRecoveryTime * time.Duration(recovery))

	if u.Stamina > EventSeasonStaminaMax {
		u.Stamina = EventSeasonStaminaMax
		checkTime = now
	}

	u.CheckedAt = DateTime{Time: checkTime}

	return true
}
