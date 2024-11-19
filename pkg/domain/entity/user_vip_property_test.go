package entity

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserVipProperty(t *testing.T) {
	type args struct {
		mirrativID string
	}
	tests := []struct {
		name string
		args args
		want *UserVipProperty
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				mirrativID: "user",
			},
			want: &UserVipProperty{
				MirrativID: "user",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserVipProperty(tt.args.mirrativID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserVipProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserVipProperty_AddExp(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		MirrativID string
		Exp        uint
		Coin       uint
	}
	type args struct {
		exp uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserVipProperty
	}{
		{
			name: "経験値加算",
			fields: fields{
				Exp: 100,
			},
			args: args{
				exp: 150,
			},
			want: &UserVipProperty{
				Exp: 250,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserVipProperty{
				Model:      tt.fields.Model,
				MirrativID: tt.fields.MirrativID,
				Exp:        tt.fields.Exp,
				Coin:       tt.fields.Coin,
			}
			u.AddExp(tt.args.exp)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("UserVipProperty.AddExp() = %v, want %v", u, tt.want)
			}
		})
	}
}

func TestUserVipProperty_AddCoin(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		MirrativID string
		Exp        uint
		Coin       uint
	}
	type args struct {
		coin uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserVipProperty
	}{
		{
			name: "コイン加算",
			fields: fields{
				Coin: 100,
			},
			args: args{
				coin: 150,
			},
			want: &UserVipProperty{
				Coin: 250,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserVipProperty{
				Model:      tt.fields.Model,
				MirrativID: tt.fields.MirrativID,
				Exp:        tt.fields.Exp,
				Coin:       tt.fields.Coin,
			}
			u.AddCoin(tt.args.coin)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("UserVipProperty.AddCoin() = %v, want %v", u, tt.want)
			}
		})
	}
}
