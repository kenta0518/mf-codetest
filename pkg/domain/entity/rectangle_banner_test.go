package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestRectangleBanner_Update(t *testing.T) {

	type fields struct {
		Model          gorm.Model
		Term           Term
		DisplayOrder   uint
		BannerID       string
		TransitionType string
	}
	type args struct {
		displayOrder    uint
		bannerID        string
		transitionType  string
		transitionValue uint
		term            Term
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RectangleBanner
	}{
		{
			name: "更新",
			fields: fields{
				Model:          gorm.Model{ID: 1},
				Term:           Term{},
				DisplayOrder:   1,
				BannerID:       "b1",
				TransitionType: "test",
			},
			args: args{
				term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
				},
				displayOrder:    2,
				bannerID:        "b2",
				transitionType:  "test2",
				transitionValue: 1,
			},
			want: &RectangleBanner{
				Model: gorm.Model{ID: 1},
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
				},
				DisplayOrder:    2,
				BannerID:        "b2",
				TransitionType:  "test2",
				TransitionValue: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &RectangleBanner{
				Model:          tt.fields.Model,
				Term:           tt.fields.Term,
				DisplayOrder:   tt.fields.DisplayOrder,
				BannerID:       tt.fields.BannerID,
				TransitionType: tt.fields.TransitionType,
			}
			b.Update(tt.args.displayOrder, tt.args.bannerID, tt.args.transitionType, tt.args.transitionValue, tt.args.term)
			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("RectangleBanner.Update() = %v, want %v", b, tt.want)
			}
		})
	}
}
