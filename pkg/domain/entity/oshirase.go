package entity

import (
	"time"

	"gorm.io/gorm"
)

const (
	Info      = iota // お知らせ
	InfoGacha        //ガチャ
	Campaign         //キャンペーン
	Events           // その他
	Bugs             // 不具合
	Important        // 重要
	Emomo            //エモモ
)

type Oshirase struct {
	gorm.Model
	Term
	Title                  string
	Message                string
	Kind                   uint
	DisplayOrder           uint
	BannerImageID          string
	AlternativePublishedAt time.Time
	PostscriptCount        uint
	PostscriptAt           *time.Time
}

func (o *Oshirase) Update(title string, message string, kind uint, displayOrder uint, BannerImageID string, alternativePublishAt time.Time,
	term Term, postscriptCount uint, postscriptAt *time.Time) {
	o.Title = title
	o.Message = message
	o.Kind = kind
	o.DisplayOrder = displayOrder
	o.BannerImageID = BannerImageID
	o.AlternativePublishedAt = alternativePublishAt
	o.Term = term
	if postscriptAt != nil {
		o.PostscriptCount = postscriptCount
		o.PostscriptAt = postscriptAt
	}
}
