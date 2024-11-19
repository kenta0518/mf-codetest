package entity

type EventAdv struct {
	SeedBase      `yaml:",inline"`
	Term          `yaml:",inline"`
	EventSeasonID uint      `yaml:"eventSeasonId"`
	StoryName     string    `yaml:"storyName"`
	UnLockStartAt *DateTime `yaml:"unlockStartAt"`
}
