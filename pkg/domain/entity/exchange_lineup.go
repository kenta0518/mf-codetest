package entity

type SalesFormKind string

const (
	SALES_FORM_KIND_NONE   SalesFormKind = "None"
	SALES_FORM_KIND_DAILY  SalesFormKind = "Daily"
	SALES_FORM_KIND_PERIOD SalesFormKind = "Period"
)

type ExchangeLineup struct {
	SeedBase      `yaml:",inline"`
	Term          `yaml:",inline"`
	ExchangeID    uint          `yaml:"groupId"`
	Name          string        `yaml:"name"`
	ThumbnailType ContentType   `yaml:"thumbnailType"`
	ThumbnailID   uint          `yaml:"thumbnailId"`
	Price         int           `yaml:"medal"`
	SalesLimit    int           `yaml:"salesLimit"`
	Ribbon        string        `yaml:"ribbon"`
	RibbonColor   string        `yaml:"ribbonColor"`
	SalesFormKind SalesFormKind `yaml:"groupKind"`
	Contents      []ExchangeContent
}
