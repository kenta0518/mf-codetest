package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewUserVip(t *testing.T) {
	type args struct {
		userID uint
		vip    *Vip
	}
	tests := []struct {
		name string
		args args
		want *UserVip
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userID: 1,
				vip:    &Vip{SeedBase: SeedBase{ID: 1000}},
			},
			want: &UserVip{
				UserID: 1,
				VipID:  1000,
				Vip:    Vip{SeedBase: SeedBase{ID: 1000}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserVip(tt.args.userID, tt.args.vip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserVip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVip_Receive(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		UserID     uint
		VipID      uint
		Vip        Vip
		ReceivedAt *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserVip
	}{
		{
			name: "受け取れるか",
			fields: fields{
				ReceivedAt: nil,
			},
			args: args{
				time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
			want: &UserVip{
				ReceivedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserVip{
				Model:      tt.fields.Model,
				UserID:     tt.fields.UserID,
				VipID:      tt.fields.VipID,
				Vip:        tt.fields.Vip,
				ReceivedAt: tt.fields.ReceivedAt,
			}
			u.Receive(tt.args.time)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("UserVip.Receive() = %v, want %v", u, tt.want)
			}
		})
	}
}

func TestUserVip_IsReceived(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		UserID     uint
		VipID      uint
		Vip        Vip
		ReceivedAt *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "受け取ってればtrue",
			fields: fields{
				ReceivedAt: func(t time.Time) *time.Time { return &t }(time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)),
			},
			want: true,
		},
		{
			name: "受け取ってなければfalse",
			fields: fields{
				ReceivedAt: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserVip{
				Model:      tt.fields.Model,
				UserID:     tt.fields.UserID,
				VipID:      tt.fields.VipID,
				Vip:        tt.fields.Vip,
				ReceivedAt: tt.fields.ReceivedAt,
			}
			if got := u.IsReceived(); got != tt.want {
				t.Errorf("UserVip.IsReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVip_IsAchieved(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		UserID     uint
		VipID      uint
		Vip        Vip
		ReceivedAt *time.Time
	}
	type args struct {
		exp uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "達成してればtrue1",
			fields: fields{
				Vip: Vip{Goal: 1000},
			},
			args: args{
				exp: 1000,
			},
			want: true,
		},
		{
			name: "達成してればtrue2",
			fields: fields{
				Vip: Vip{Goal: 1000},
			},
			args: args{
				exp: 1000,
			},
			want: true,
		},
		{
			name: "達成してなければfalse",
			fields: fields{
				Vip: Vip{Goal: 1000},
			},
			args: args{
				exp: 999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserVip{
				Model:      tt.fields.Model,
				UserID:     tt.fields.UserID,
				VipID:      tt.fields.VipID,
				Vip:        tt.fields.Vip,
				ReceivedAt: tt.fields.ReceivedAt,
			}
			if got := u.IsAchieved(tt.args.exp); got != tt.want {
				t.Errorf("UserVip.IsAchieved() = %v, want %v", got, tt.want)
			}
		})
	}
}
