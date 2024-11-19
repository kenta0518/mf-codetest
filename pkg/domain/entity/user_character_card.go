package entity

const (
	InitCharacterCardID uint = 101101
)

type UserCharacterCard struct {
	UserResourceBase
	CharacterCard CharacterCard `gorm:"foreignKey:ResourceID"`
	Level         uint          `gorm:"not null"`
	Awake         uint          `gorm:"not null"`
}

func NewUserCharacterCard(userID, characterCardID uint) *UserCharacterCard {
	return &UserCharacterCard{
		UserResourceBase: UserResourceBase{
			UserID:     userID,
			ResourceID: characterCardID,
		},
		Level: 1,
		Awake: 0,
	}
}

func (u *UserCharacterCard) AddLevel(level uint) {
	u.Level += level
}

func (u *UserCharacterCard) UpAwake() {
	u.Awake++
}
