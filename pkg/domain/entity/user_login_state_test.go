package entity

import (
	"reflect"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"gorm.io/gorm"
)

func TestUserLoginState_HasLoggedInToday(t *testing.T) {
	lastLoginAt1 := time.Date(2023, 5, 1, 15, 0, 0, 0, time.Local)
	lastLoginAt2 := time.Date(2023, 4, 30, 7, 0, 0, 0, time.Local)
	flextime.Fix(time.Date(2023, 5, 1, 22, 0, 0, 0, time.Local))

	type fields struct {
		Model         gorm.Model
		UserID        uint
		Total         uint
		Duration      *uint
		DurationStart *time.Time
		LastLoginAt   *time.Time
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "今日ログインしたか",
			fields: fields{
				LastLoginAt: &lastLoginAt1,
			},
			args: args{
				user: &User{},
			},
			want: true,
		},
		{
			name: "今日ログインしてない1",
			fields: fields{
				LastLoginAt: nil,
			},
			args: args{
				user: &User{},
			},
			want: false,
		},
		{
			name: "今日ログインしてない2",
			fields: fields{
				LastLoginAt: &lastLoginAt2,
			},
			args: args{
				user: &User{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserLoginState{
				Model:         tt.fields.Model,
				UserID:        tt.fields.UserID,
				Total:         tt.fields.Total,
				Duration:      tt.fields.Duration,
				DurationStart: tt.fields.DurationStart,
				LastLoginAt:   tt.fields.LastLoginAt,
			}
			if got := u.HasLoggedInToday(tt.args.user); got != tt.want {
				t.Errorf("UserLoginState.HasLoggedInToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLoginState_Login(t *testing.T) {
	now := time.Date(2023, 5, 1, 22, 0, 0, 0, time.Local)
	LastLoginAt1 := time.Date(2023, 4, 30, 15, 0, 0, 0, time.Local)
	LastLoginAt2 := time.Date(2023, 4, 29, 19, 0, 0, 0, time.Local)

	type fields struct {
		Model         gorm.Model
		UserID        uint
		Total         uint
		Duration      *uint
		DurationStart *time.Time
		LastLoginAt   *time.Time
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *UserLoginState
	}{
		{
			name: "初ログイン",
			fields: fields{
				Duration:    func(d uint) *uint { return &d }(0),
				LastLoginAt: nil,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLoginState{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				DurationStart: &now,
				LastLoginAt:   &now,
			},
		},
		{
			name: "連続ログイン",
			fields: fields{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				LastLoginAt:   &LastLoginAt1,
				DurationStart: &LastLoginAt1,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLoginState{
				Total:         2,
				Duration:      func(d uint) *uint { return &d }(2),
				DurationStart: &LastLoginAt1,
				LastLoginAt:   &now,
			},
		},
		{
			name: "久しぶりログイン",
			fields: fields{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				LastLoginAt:   &LastLoginAt2,
				DurationStart: &LastLoginAt2,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLoginState{
				Total:         2,
				Duration:      func(d uint) *uint { return &d }(1),
				DurationStart: &now,
				LastLoginAt:   &now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flextime.Fix(now)

			u := &UserLoginState{
				Model:         tt.fields.Model,
				UserID:        tt.fields.UserID,
				Total:         tt.fields.Total,
				Duration:      tt.fields.Duration,
				DurationStart: tt.fields.DurationStart,
				LastLoginAt:   tt.fields.LastLoginAt,
			}

			if got := u.Login(tt.args.user); got != tt.want {
				t.Errorf("UserLoginState.Login() = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual([]any{*u.Duration, *u.DurationStart, *u.LastLoginAt}, []any{*tt.want1.Duration, tt.want1.DurationStart, tt.want1.LastLoginAt}) {
				t.Errorf("UserLoginState.Login() = %v, want %v", u, tt.want1)
			}
		})
	}
}
