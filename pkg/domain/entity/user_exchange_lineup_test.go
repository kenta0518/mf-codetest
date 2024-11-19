package entity

import (
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestUserExchangeLineup_ConsumeStock(t *testing.T) {
	type fields struct {
		Model            gorm.Model
		UserID           uint
		ExchangeLineupID uint
		ExchangeLineup   ExchangeLineup
		Stock            int
		ResetAt          *time.Time
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "消費できるか",
			fields: fields{
				Stock:          10,
				ExchangeLineup: ExchangeLineup{SalesLimit: 10},
			},
			args: args{
				quantity: 5,
			},
			want: true,
		},
		{
			name: "足りない時は消費できない",
			fields: fields{
				Stock:          10,
				ExchangeLineup: ExchangeLineup{SalesLimit: 10},
			},
			args: args{
				quantity: 15,
			},
			want: false,
		},
		{
			name: "無制限の時は消費できる",
			fields: fields{
				Stock:          0,
				ExchangeLineup: ExchangeLineup{SalesLimit: 0},
			},
			args: args{
				quantity: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserExchangeLineup{
				Model:            tt.fields.Model,
				UserID:           tt.fields.UserID,
				ExchangeLineupID: tt.fields.ExchangeLineupID,
				ExchangeLineup:   tt.fields.ExchangeLineup,
				Stock:            tt.fields.Stock,
				ResetAt:          tt.fields.ResetAt,
			}
			if got := u.ConsumeStock(tt.args.quantity); got != tt.want {
				t.Errorf("UserExchangeLineup.ConsumeStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserExchangeLineup_Reset(t *testing.T) {
	type fields struct {
		Model            gorm.Model
		UserID           uint
		ExchangeLineupID uint
		ExchangeLineup   ExchangeLineup
		Stock            int
		ResetAt          *time.Time
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *UserExchangeLineup
	}{
		{
			name: "リセットできるか",
			fields: fields{
				Stock:          0,
				ExchangeLineup: ExchangeLineup{SalesLimit: 10},
				ResetAt:        nil,
			},
			args: args{
				t: time.Date(2023, 9, 1, 0, 0, 0, 0, time.Local),
			},
			want: true,
			want1: &UserExchangeLineup{
				Stock:          10,
				ResetAt:        func(t time.Time) *time.Time { return &t }(time.Date(2023, 9, 1, 0, 0, 0, 0, time.Local)),
				ExchangeLineup: ExchangeLineup{SalesLimit: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserExchangeLineup{
				Model:            tt.fields.Model,
				UserID:           tt.fields.UserID,
				ExchangeLineupID: tt.fields.ExchangeLineupID,
				ExchangeLineup:   tt.fields.ExchangeLineup,
				Stock:            tt.fields.Stock,
				ResetAt:          tt.fields.ResetAt,
			}
			if got := u.Reset(tt.args.t); got != tt.want {
				t.Errorf("UserExchangeLineup.Reset() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(u, tt.want1) {
				t.Errorf("UserExchangeLineup.Reset() = %v, want1 %v", u, tt.want1)
			}
		})
	}
}

func TestNewUserExchangeLineup(t *testing.T) {
	type args struct {
		userId         uint
		exchangeLineup *ExchangeLineup
	}
	tests := []struct {
		name string
		args args
		want *UserExchangeLineup
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userId:         1,
				exchangeLineup: &ExchangeLineup{SeedBase: SeedBase{ID: 1001}, SalesLimit: 10},
			},
			want: &UserExchangeLineup{
				UserID:           1,
				ExchangeLineupID: 1001,
				ExchangeLineup:   ExchangeLineup{SeedBase: SeedBase{ID: 1001}, SalesLimit: 10},
				Stock:            10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserExchangeLineup(tt.args.userId, tt.args.exchangeLineup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserExchangeLineup() = %v, want %v", got, tt.want)
			}
		})
	}
}
