package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	AssetBundle  = "assetbundle"
	ClientMaster = "clientmaster"
)

var ResourceCategories = []string{AssetBundle, ClientMaster}

type ResourceVersion struct {
	gorm.Model
	Category    string
	VersionCode string
	StartAt     time.Time
	EndAt       time.Time
}

func NewResouseVersion(category string, versionCode string, startAt time.Time, endAt time.Time) *ResourceVersion {
	return &ResourceVersion{
		Category:    category,
		VersionCode: versionCode,
		StartAt:     startAt,
		EndAt:       endAt,
	}
}

func (r *ResourceVersion) Edit(category string, versioncode string, startAt time.Time, endAt time.Time) {
	r.Category = category
	r.VersionCode = versioncode
	r.StartAt = startAt
	r.EndAt = endAt
}
