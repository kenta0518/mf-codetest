package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewUserInterstitialBanner(t *testing.T) {
	type args struct {
		userID uint
		banner *InterstitialBanner
	}
	tests := []struct {
		name string
		args args
		want *UserInterstitialBanner
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userID: 1,
				banner: &InterstitialBanner{
					Model: gorm.Model{ID: 1000},
				},
			},
			want: &UserInterstitialBanner{
				UserID:               1,
				InterstitialBannerID: 1000,
				InterstitialBanner: InterstitialBanner{
					Model: gorm.Model{ID: 1000},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserInterstitialBanner(tt.args.userID, tt.args.banner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserInterstitialBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserInterstitialBanner_View(t *testing.T) {
	type fields struct {
		Model                gorm.Model
		UserID               uint
		User                 User
		InterstitialBannerID uint
		InterstitialBanner   InterstitialBanner
		ViewCount            uint
		LastViewedAt         *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserInterstitialBanner
	}{
		{
			name: "インステみた",
			fields: fields{
				ViewCount:    0,
				LastViewedAt: nil,
			},
			args: args{
				time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
			want: &UserInterstitialBanner{
				ViewCount:    1,
				LastViewedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserInterstitialBanner{
				Model:                tt.fields.Model,
				UserID:               tt.fields.UserID,
				User:                 tt.fields.User,
				InterstitialBannerID: tt.fields.InterstitialBannerID,
				InterstitialBanner:   tt.fields.InterstitialBanner,
				ViewCount:            tt.fields.ViewCount,
				LastViewedAt:         tt.fields.LastViewedAt,
			}
			u.View(tt.args.time)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("UserInterstitialBanner.View() = %v, want %v", u, tt.want)
			}
		})
	}
}

func TestUserInterstitialBanner_IsDisplayable(t *testing.T) {
	type fields struct {
		Model                gorm.Model
		UserID               uint
		User                 User
		InterstitialBannerID uint
		InterstitialBanner   InterstitialBanner
		ViewCount            uint
		LastViewedAt         *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "表示可能1",
			fields: fields{
				ViewCount:    0,
				LastViewedAt: nil,
				InterstitialBanner: InterstitialBanner{
					Span:      1 * time.Hour,
					ViewCount: 2,
				},
			},
			args: args{
				time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "表示可能2",
			fields: fields{
				ViewCount:    1,
				LastViewedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
				InterstitialBanner: InterstitialBanner{
					Span:      1 * time.Hour,
					ViewCount: 2,
				},
			},
			args: args{
				time: time.Date(2023, 5, 1, 1, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "表示不可能1",
			fields: fields{
				ViewCount:    2,
				LastViewedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
				InterstitialBanner: InterstitialBanner{
					Span:      1 * time.Hour,
					ViewCount: 2,
				},
			},
			args: args{
				time: time.Date(2023, 5, 1, 1, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "表示不可能2",
			fields: fields{
				ViewCount:    1,
				LastViewedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
				InterstitialBanner: InterstitialBanner{
					Span:      1 * time.Hour,
					ViewCount: 2,
				},
			},
			args: args{
				time: time.Date(2023, 5, 1, 0, 30, 0, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserInterstitialBanner{
				Model:                tt.fields.Model,
				UserID:               tt.fields.UserID,
				User:                 tt.fields.User,
				InterstitialBannerID: tt.fields.InterstitialBannerID,
				InterstitialBanner:   tt.fields.InterstitialBanner,
				ViewCount:            tt.fields.ViewCount,
				LastViewedAt:         tt.fields.LastViewedAt,
			}
			if got := u.IsDisplayable(tt.args.time); got != tt.want {
				t.Errorf("UserInterstitialBanner.IsDisplayable() = %v, want %v", got, tt.want)
			}
		})
	}
}
