package entity

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserPrizeSticker(t *testing.T) {
	type args struct {
		userID         uint
		prizeStickerID uint
		quantity       int
	}
	tests := []struct {
		name string
		args args
		want *UserPrizeSticker
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userID:         1,
				prizeStickerID: 1001,
				quantity:       1,
			},
			want: &UserPrizeSticker{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1001,
				},
				Quantity: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserPrizeSticker(tt.args.userID, tt.args.prizeStickerID, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserPrizeSticker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPrizeSticker_Gain(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		PrizeSticker     PrizeSticker
		Quantity         int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *UserPrizeSticker
	}{
		{
			name: "獲得",
			fields: fields{
				Quantity: 0,
			},
			args: args{
				quantity: 1,
			},
			want: true,
			want1: &UserPrizeSticker{
				Quantity: 1,
			},
		},
		{
			name: "獲得失敗",
			fields: fields{
				Quantity: 0,
			},
			args: args{
				quantity: -1,
			},
			want: false,
			want1: &UserPrizeSticker{
				Quantity: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPrizeSticker{
				UserResourceBase: tt.fields.UserResourceBase,
				PrizeSticker:     tt.fields.PrizeSticker,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.Gain(tt.args.quantity); got != tt.want {
				t.Errorf("UserPrizeSticker.Gain() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(u, tt.want1) {
				t.Errorf("UserPrizeSticker.Gain() = %v, want1 %v", u, tt.want1)
			}
		})
	}
}

func TestUserPrizeSticker_Consume(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		PrizeSticker     PrizeSticker
		Quantity         int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  *UserPrizeSticker
	}{
		{
			name: "消費できるか",
			fields: fields{
				Quantity: 10,
			},
			args: args{
				quantity: 5,
			},
			want: true,
			want1: &UserPrizeSticker{
				Quantity: 5,
			},
		},
		{
			name: "足りない時は消費できない",
			fields: fields{
				Quantity: 10,
			},
			args: args{
				quantity: 15,
			},
			want: false,
			want1: &UserPrizeSticker{
				Quantity: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPrizeSticker{
				UserResourceBase: tt.fields.UserResourceBase,
				PrizeSticker:     tt.fields.PrizeSticker,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.Consume(tt.args.quantity); got != tt.want {
				t.Errorf("UserPrizeSticker.Consume() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(u, tt.want1) {
				t.Errorf("UserPrizeSticker.Consume() = %v, want1 %v", u, tt.want1)
			}
		})
	}
}

func TestUserPrizeSticker_IsNew(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		PrizeSticker     PrizeSticker
		Quantity         int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IDが0なら新規",
			fields: fields{
				UserResourceBase: UserResourceBase{
					Model: gorm.Model{
						ID: 0,
					},
				},
			},
			want: true,
		},
		{
			name: "IDが0以外なら新規ではない",
			fields: fields{
				UserResourceBase: UserResourceBase{
					Model: gorm.Model{
						ID: 1,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserPrizeSticker{
				UserResourceBase: tt.fields.UserResourceBase,
				PrizeSticker:     tt.fields.PrizeSticker,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.IsEmpty(); got != tt.want {
				t.Errorf("UserPrizeSticker.IsNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
