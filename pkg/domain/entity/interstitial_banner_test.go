package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestInterstitialBanner_Update(t *testing.T) {
	type fields struct {
		Model           gorm.Model
		Term            Term
		Image           string
		Span            time.Duration
		ViewCount       uint
		TransitionType  string
		TransitionValue int
	}
	type args struct {
		img     string
		span    time.Duration
		vc      uint
		txType  string
		txValue int
		term    Term
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InterstitialBanner
	}{
		{
			name: "更新",
			fields: fields{
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local)},
				},
				Image:           "test1",
				Span:            time.Duration(1),
				ViewCount:       1,
				TransitionType:  "test1",
				TransitionValue: 1,
			},
			args: args{
				img:     "test2",
				span:    time.Duration(2),
				vc:      2,
				txType:  "test2",
				txValue: 2,
				term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)},
				},
			},
			want: &InterstitialBanner{
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2024, 6, 1, 0, 0, 0, 0, time.Local)},
				},
				Image:           "test2",
				Span:            time.Duration(2),
				ViewCount:       2,
				TransitionType:  "test2",
				TransitionValue: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &InterstitialBanner{
				Model:           tt.fields.Model,
				Term:            tt.fields.Term,
				Image:           tt.fields.Image,
				Span:            tt.fields.Span,
				ViewCount:       tt.fields.ViewCount,
				TransitionType:  tt.fields.TransitionType,
				TransitionValue: tt.fields.TransitionValue,
			}
			b.Update(tt.args.img, tt.args.span, tt.args.vc, tt.args.txType, tt.args.txValue, tt.args.term)
			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("InterstitialBanner.Update() = %v, want %v", b, tt.want)
			}
		})
	}
}
