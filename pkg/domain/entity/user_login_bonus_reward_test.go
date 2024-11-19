package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewUserLoginBonusReward(t *testing.T) {
	type args struct {
		userID   uint
		rewardID uint
	}
	tests := []struct {
		name string
		args args
		want *UserLoginBonusReward
	}{
		{
			name: "インスタンス生成できるか",
			args: args{
				userID:   1000,
				rewardID: 100,
			},
			want: &UserLoginBonusReward{
				UserID:             1000,
				LoginBonusRewardID: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserLoginBonusReward(tt.args.userID, tt.args.rewardID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserLoginBonusReward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLoginBonusReward_IsReceived(t *testing.T) {
	type fields struct {
		Model              gorm.Model
		UserID             uint
		LoginBonusRewardID uint
		LoginBonusReward   LoginBonusReward
		ReceivedAt         *time.Time
		ResetAt            *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "受け取ってたらtrue",
			fields: fields{
				ReceivedAt: &time.Time{},
			},
			want: true,
		},
		{
			name: "受け取ってたらfalse",
			fields: fields{
				ReceivedAt: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserLoginBonusReward{
				Model:              tt.fields.Model,
				UserID:             tt.fields.UserID,
				LoginBonusRewardID: tt.fields.LoginBonusRewardID,
				LoginBonusReward:   tt.fields.LoginBonusReward,
				ReceivedAt:         tt.fields.ReceivedAt,
				ResetAt:            tt.fields.ResetAt,
			}
			if got := u.IsReceived(); got != tt.want {
				t.Errorf("UserLoginBonusReward.IsReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLoginBonusReward_Receive(t *testing.T) {
	type fields struct {
		Model              gorm.Model
		UserID             uint
		LoginBonusRewardID uint
		LoginBonusReward   LoginBonusReward
		ReceivedAt         *time.Time
		ResetAt            *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserLoginBonusReward
	}{
		{
			name: "受取り",
			fields: fields{
				ReceivedAt: nil,
			},
			args: args{
				time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
			want: &UserLoginBonusReward{
				ReceivedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLoginBonusReward{
				Model:              tt.fields.Model,
				UserID:             tt.fields.UserID,
				LoginBonusRewardID: tt.fields.LoginBonusRewardID,
				LoginBonusReward:   tt.fields.LoginBonusReward,
				ReceivedAt:         tt.fields.ReceivedAt,
				ResetAt:            tt.fields.ResetAt,
			}
			u.Receive(tt.args.time)
		})
	}
}

func TestUserLoginBonusReward_Reset(t *testing.T) {
	type fields struct {
		Model              gorm.Model
		UserID             uint
		LoginBonusRewardID uint
		LoginBonusReward   LoginBonusReward
		ReceivedAt         *time.Time
		ResetAt            *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserLoginBonusReward
	}{
		{
			name: "リセット",
			fields: fields{
				ReceivedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
			},
			args: args{
				time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: &UserLoginBonusReward{
				ReceivedAt: nil,
				ResetAt:    func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLoginBonusReward{
				Model:              tt.fields.Model,
				UserID:             tt.fields.UserID,
				LoginBonusRewardID: tt.fields.LoginBonusRewardID,
				LoginBonusReward:   tt.fields.LoginBonusReward,
				ReceivedAt:         tt.fields.ReceivedAt,
				ResetAt:            tt.fields.ResetAt,
			}
			u.Reset(tt.args.time)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("UserLoginBonusReward.Reset() = %v, want %v", u, tt.want)
			}
		})
	}
}
