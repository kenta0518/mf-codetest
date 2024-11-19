package entity

import "gorm.io/gorm"

const (
	DeckSize = 4
)

type SlotIndex uint
type SlotType string

const (
	SlotIndex1 SlotIndex = iota
	SlotIndex2
	SlotIndex3
	SlotIndex4
)

const (
	SlotTypeMain SlotType = "main"
	SlotTypeSub  SlotType = "sub"
)

type UserDeck struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`

	MainUserCharacterCardID1 uint `gorm:"not null"` // Main1は必須
	MainUserCharacterCardID2 uint
	MainUserCharacterCardID3 uint
	MainUserCharacterCardID4 uint

	SubUserCharacterCardID1 uint
	SubUserCharacterCardID2 uint
	SubUserCharacterCardID3 uint
	SubUserCharacterCardID4 uint
}

func NewUserDeck(userID, characterCardID1 uint) *UserDeck {
	return &UserDeck{
		UserID:                   userID,
		MainUserCharacterCardID1: characterCardID1,
	}
}

func (d *UserDeck) SafeChange(slotIndex SlotIndex, slotType SlotType, charaCardID uint) bool {
	// SlotIndexの範囲チェック
	if slotIndex < SlotIndex1 || slotIndex > SlotIndex4 {
		return false
	}
	// SlotIndeのMainは必須
	if slotIndex == SlotIndex1 && slotType == SlotTypeMain && charaCardID == 0 {
		return false
	}

	tempCardID := uint(0)
	if slotType == SlotTypeMain {
		tempCardID = d.getMainCharaCardID(slotIndex)
	} else {
		tempCardID = d.getSubCharaCardID(slotIndex)
	}

	if tempCardID != charaCardID && charaCardID != 0 {
		d.checkOrganizedCharaCard(charaCardID, tempCardID)
	}

	if slotType == SlotTypeMain {
		d.updateMainCharaCardID(slotIndex, charaCardID)
	} else {
		d.updateSubCharaCardID(slotIndex, charaCardID)
	}

	return true
}

func (d *UserDeck) checkOrganizedCharaCard(targetID, changeID uint) {
	if d.MainUserCharacterCardID1 == targetID {
		d.MainUserCharacterCardID1 = changeID
	}
	if d.MainUserCharacterCardID2 == targetID {
		d.MainUserCharacterCardID2 = changeID
	}
	if d.MainUserCharacterCardID3 == targetID {
		d.MainUserCharacterCardID3 = changeID
	}
	if d.MainUserCharacterCardID4 == targetID {
		d.MainUserCharacterCardID4 = changeID
	}

	if d.SubUserCharacterCardID1 == targetID {
		d.SubUserCharacterCardID1 = changeID
	}
	if d.SubUserCharacterCardID2 == targetID {
		d.SubUserCharacterCardID2 = changeID
	}
	if d.SubUserCharacterCardID3 == targetID {
		d.SubUserCharacterCardID3 = changeID
	}
	if d.SubUserCharacterCardID4 == targetID {
		d.SubUserCharacterCardID4 = changeID
	}
}

func (d *UserDeck) getMainCharaCardID(slotIndex SlotIndex) uint {
	switch slotIndex {
	case SlotIndex1:
		return d.MainUserCharacterCardID1
	case SlotIndex2:
		return d.MainUserCharacterCardID2
	case SlotIndex3:
		return d.MainUserCharacterCardID3
	case SlotIndex4:
		return d.MainUserCharacterCardID4
	}
	return 0
}

func (d *UserDeck) getSubCharaCardID(slotIndex SlotIndex) uint {
	switch slotIndex {
	case SlotIndex1:
		return d.SubUserCharacterCardID1
	case SlotIndex2:
		return d.SubUserCharacterCardID2
	case SlotIndex3:
		return d.SubUserCharacterCardID3
	case SlotIndex4:
		return d.SubUserCharacterCardID4
	}
	return 0
}

func (d *UserDeck) updateMainCharaCardID(slotIndex SlotIndex, charaCardID uint) {
	switch slotIndex {
	case SlotIndex1:
		d.MainUserCharacterCardID1 = charaCardID
	case SlotIndex2:
		d.MainUserCharacterCardID2 = charaCardID
	case SlotIndex3:
		d.MainUserCharacterCardID3 = charaCardID
	case SlotIndex4:
		d.MainUserCharacterCardID4 = charaCardID
	}
}

func (d *UserDeck) updateSubCharaCardID(slotIndex SlotIndex, charaCardID uint) {
	switch slotIndex {
	case SlotIndex1:
		d.SubUserCharacterCardID1 = charaCardID
	case SlotIndex2:
		d.SubUserCharacterCardID2 = charaCardID
	case SlotIndex3:
		d.SubUserCharacterCardID3 = charaCardID
	case SlotIndex4:
		d.SubUserCharacterCardID4 = charaCardID
	}
}
