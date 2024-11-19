package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewNewsHeadline(t *testing.T) {
	type args struct {
		text        string
		startAt     time.Time
		endAt       time.Time
		testStartAt time.Time
		testEndAt   time.Time
		order       uint
	}
	tests := []struct {
		name string
		args args
		want *NewsHeadline
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				text:        "text",
				startAt:     time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
				endAt:       time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local),
				testStartAt: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
				testEndAt:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local),
				order:       1,
			},
			want: &NewsHeadline{
				Text: "text",
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
				},
				DisplayOrder: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNewsHeadline(tt.args.text, tt.args.startAt, tt.args.endAt, tt.args.testStartAt, tt.args.testEndAt, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNewsHeadline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewsHeadline_Update(t *testing.T) {
	type fields struct {
		Model        gorm.Model
		Term         Term
		Text         string
		DisplayOrder uint
	}
	type args struct {
		text         string
		startAt      time.Time
		endAt        time.Time
		testStartAt  time.Time
		testEndAt    time.Time
		displayOrder uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *NewsHeadline
	}{
		{
			name: "更新",
			fields: fields{
				Text: "text1",
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local)},
				},
				DisplayOrder: 1,
			},
			args: args{
				text:         "text2",
				startAt:      time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
				endAt:        time.Date(2023, 6, 2, 0, 0, 0, 0, time.Local),
				testStartAt:  time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
				testEndAt:    time.Date(2023, 6, 2, 0, 0, 0, 0, time.Local),
				displayOrder: 1,
			},
			want: &NewsHeadline{
				Text: "text2",
				Term: Term{
					StartAt:     DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
					EndAt:       DateTime{Time: time.Date(2023, 6, 2, 0, 0, 0, 0, time.Local)},
					TestStartAt: DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
					TestEndAt:   DateTime{Time: time.Date(2023, 6, 2, 0, 0, 0, 0, time.Local)},
				},
				DisplayOrder: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NewsHeadline{
				Model:        tt.fields.Model,
				Term:         tt.fields.Term,
				Text:         tt.fields.Text,
				DisplayOrder: tt.fields.DisplayOrder,
			}
			n.Update(tt.args.text, tt.args.startAt, tt.args.endAt, tt.args.testStartAt, tt.args.testEndAt, tt.args.displayOrder)
			if !reflect.DeepEqual(n, tt.want) {
				t.Errorf("NewsHeadline.Update() = %v, want %v", n, tt.want)
			}
		})
	}
}
