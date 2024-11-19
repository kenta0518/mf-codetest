package entity

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestOshirase_Update(t *testing.T) {
	tm := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	type fields struct {
		Model                  gorm.Model
		Term                   Term
		Title                  string
		Message                string
		Kind                   uint
		DisplayOrder           uint
		BannerImageID          string
		AlternativePublishedAt time.Time
		PostscriptCount        uint
		PostscriptAt           *time.Time
	}
	type args struct {
		title                string
		message              string
		kind                 uint
		displayOrder         uint
		bannerImageId        string
		alternativePublishAt time.Time
		term                 Term
		postscriptCount      uint
		postscriptAt         *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Oshirase
	}{
		{
			name: "お知らせ更新",
			fields: fields{
				Model:                  gorm.Model{ID: 1000},
				Title:                  "お知らせタイトル",
				Message:                "お知らせメッセージ",
				Kind:                   1,
				DisplayOrder:           1,
				BannerImageID:          "1",
				AlternativePublishedAt: time.Now(),
				PostscriptCount:        1,
				PostscriptAt:           &time.Time{},
				Term:                   *NewTerm(time.Now(), time.Now(), time.Now(), time.Now()),
			},
			args: args{
				title:                "お知らせタイトル更新",
				message:              "お知らせメッセージ更新",
				kind:                 2,
				displayOrder:         2,
				bannerImageId:        "1",
				alternativePublishAt: tm,
				term:                 *NewTerm(tm, tm, tm, tm),
				postscriptCount:      2,
				postscriptAt:         &tm,
			},
			want: &Oshirase{
				Model:                  gorm.Model{ID: 1000},
				Title:                  "お知らせタイトル更新",
				Message:                "お知らせメッセージ更新",
				Kind:                   2,
				DisplayOrder:           2,
				BannerImageID:          "1",
				AlternativePublishedAt: tm,
				PostscriptCount:        2,
				PostscriptAt:           &tm,
				Term:                   *NewTerm(tm, tm, tm, tm),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Oshirase{
				Model:                  tt.fields.Model,
				Term:                   tt.fields.Term,
				Title:                  tt.fields.Title,
				Message:                tt.fields.Message,
				Kind:                   tt.fields.Kind,
				DisplayOrder:           tt.fields.DisplayOrder,
				BannerImageID:          tt.fields.BannerImageID,
				AlternativePublishedAt: tt.fields.AlternativePublishedAt,
				PostscriptCount:        tt.fields.PostscriptCount,
				PostscriptAt:           tt.fields.PostscriptAt,
			}
			o.Update(tt.args.title, tt.args.message, tt.args.kind, tt.args.displayOrder, tt.args.bannerImageId, tt.args.alternativePublishAt, tt.args.term, tt.args.postscriptCount, tt.args.postscriptAt)
		})
	}
}
