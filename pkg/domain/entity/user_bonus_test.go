package entity

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserBonus(t *testing.T) {
	type args struct {
		userID   uint
		bonusID  uint
		quantity int
	}
	tests := []struct {
		name string
		args args
		want *UserBonus
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userID:   1000,
				bonusID:  1000,
				quantity: 3,
			},
			want: &UserBonus{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1000,
				},
				Quantity: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserBonus(tt.args.userID, tt.args.bonusID, tt.args.quantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserBonus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserBonus_Gain(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		BonusID          uint
		Bonus            Bonus
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
		want1  *UserBonus
	}{
		{
			name: "増やせるか",
			fields: fields{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1000,
				},
				Quantity: 5,
			},
			args: args{
				quantity: 1,
			},
			want: true,
			want1: &UserBonus{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1000,
				},
				Quantity: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserBonus{
				UserResourceBase: tt.fields.UserResourceBase,
				Bonus:            tt.fields.Bonus,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.Gain(tt.args.quantity); got != tt.want {
				t.Errorf("UserBonus.Gain() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(u, tt.want1) {
				t.Errorf("UserBonus = %v, want %v", u, tt.want1)
			}
		})
	}
}

func TestUserBonus_IsNew(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		Bonus            Bonus
		Quantity         int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "新規レコードならtrue",
			fields: fields{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1000,
				},
				Quantity: 1,
			},
			want: true,
		},
		{
			name: "既存レコードならfalse",
			fields: fields{
				UserResourceBase: UserResourceBase{
					Model:      gorm.Model{ID: 1000},
					UserID:     1000,
					ResourceID: 1000,
				},
				Quantity: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserBonus{
				UserResourceBase: tt.fields.UserResourceBase,
				Bonus:            tt.fields.Bonus,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.IsEmpty(); got != tt.want {
				t.Errorf("UserBonus.IsNew() = %v, want %v", got, tt.want)
			}
		})
	}
}
