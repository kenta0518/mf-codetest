package entity

type GachaType string

const (
	GACHA_TYPE_CONSTANT GachaType = "Constant" // 恒常ガチャ
	GACHA_TYPE_LIMITED  GachaType = "Limited"  // 有償限定ガチャ
	GACHA_TYPE_RESOURCE GachaType = "Resource" // Resourceガチャ
)

type Gacha struct {
	SeedBase           `yaml:",inline"`
	Term               `yaml:",inline"`
	Name               string      `yaml:"name" gorm:"size:512"`
	Order              int         `yaml:"order"`
	PayContentType     ContentType `yaml:"payContentType"`
	PayContentId       uint        `yaml:"payContentId"`
	PayContentQuantity uint        `yaml:"payContentQuantity"`
	MiniBanner         string      `yaml:"miniBanner"`
	Banner             string      `yaml:"banner"`
	Note               string      `yaml:"note" gorm:"size:2048"`
}

type WeightPair[T any] struct {
	Target T
	Weight uint
}
