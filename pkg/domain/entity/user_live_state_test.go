package entity

import (
	"reflect"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"gorm.io/gorm"
)

func TestUserLiveState_HasLivedToday(t *testing.T) {
	lastLiveAt1 := time.Date(2023, 5, 1, 15, 0, 0, 0, time.Local)
	lastLiveAt2 := time.Date(2023, 4, 30, 7, 0, 0, 0, time.Local)
	flextime.Fix(time.Date(2023, 5, 1, 22, 0, 0, 0, time.Local))

	type fields struct {
		Model         gorm.Model
		UserID        uint
		Total         uint
		Duration      *uint
		DurationStart *time.Time
		LastLiveAt    *time.Time
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
			name: "今日配信したか",
			fields: fields{
				LastLiveAt: &lastLiveAt1,
			},
			args: args{
				user: &User{},
			},
			want: true,
		},
		{
			name: "今日配信してない1",
			fields: fields{
				LastLiveAt: nil,
			},
			args: args{
				user: &User{},
			},
			want: false,
		},
		{
			name: "今日配信してない2",
			fields: fields{
				LastLiveAt: &lastLiveAt2,
			},
			args: args{
				user: &User{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserLiveState{
				Model:         tt.fields.Model,
				UserID:        tt.fields.UserID,
				Total:         tt.fields.Total,
				Duration:      tt.fields.Duration,
				DurationStart: tt.fields.DurationStart,
				LastLiveAt:    tt.fields.LastLiveAt,
			}
			if got := u.HasLivedToday(tt.args.user); got != tt.want {
				t.Errorf("UserLiveState.HasLivedToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLiveState_Live(t *testing.T) {
	now := time.Date(2023, 5, 1, 22, 0, 0, 0, time.Local)
	LastLiveAt1 := time.Date(2023, 4, 30, 15, 0, 0, 0, time.Local)
	LastLiveAt2 := time.Date(2023, 4, 29, 19, 0, 0, 0, time.Local)

	type fields struct {
		Model         gorm.Model
		UserID        uint
		Total         uint
		Duration      *uint
		DurationStart *time.Time
		LastLiveAt    *time.Time
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *UserLiveState
	}{
		{
			name: "初配信",
			fields: fields{
				Duration:   func(d uint) *uint { return &d }(0),
				LastLiveAt: nil,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLiveState{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				DurationStart: &now,
				LastLiveAt:    &now,
			},
		},
		{
			name: "連続配信",
			fields: fields{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				LastLiveAt:    &LastLiveAt1,
				DurationStart: &LastLiveAt1,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLiveState{
				Total:         2,
				Duration:      func(d uint) *uint { return &d }(2),
				DurationStart: &LastLiveAt1,
				LastLiveAt:    &now,
			},
		},
		{
			name: "久しぶり配信",
			fields: fields{
				Total:         1,
				Duration:      func(d uint) *uint { return &d }(1),
				LastLiveAt:    &LastLiveAt2,
				DurationStart: &LastLiveAt2,
			},
			args: args{
				user: &User{},
			},
			want: true,
			want1: &UserLiveState{
				Total:         2,
				Duration:      func(d uint) *uint { return &d }(1),
				DurationStart: &now,
				LastLiveAt:    &now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLiveState{
				Model:         tt.fields.Model,
				UserID:        tt.fields.UserID,
				Total:         tt.fields.Total,
				Duration:      tt.fields.Duration,
				DurationStart: tt.fields.DurationStart,
				LastLiveAt:    tt.fields.LastLiveAt,
			}
			if got := u.Live(tt.args.user); got != tt.want {
				t.Errorf("UserLiveState.Live() = %v, want %v", got, tt.want)
			}
			if reflect.DeepEqual([]any{*u.Duration, *u.DurationStart, *u.LastLiveAt}, []any{*tt.want1.Duration, tt.want1.DurationStart, tt.want1.LastLiveAt}) {
				t.Errorf("UserLiveState.Live() = %v, want %v", u, tt.want1)
			}
		})
	}
}
