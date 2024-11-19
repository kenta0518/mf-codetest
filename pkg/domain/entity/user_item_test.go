package entity

import (
	"testing"
)

func TestUserItem_Gain(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		Item             Item
		Quantity         int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
		want   bool
	}{
		{
			name:   "アイテム数増やせるか",
			fields: fields{Quantity: 5},
			args:   args{1},
			result: 6,
			want:   true,
		},
		{
			name:   "負の数はいったら失敗",
			fields: fields{Quantity: 5},
			args:   args{-1},
			result: 5,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserItem{
				UserResourceBase: tt.fields.UserResourceBase,
				Item:             tt.fields.Item,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.Gain(tt.args.quantity); got != tt.want {
				t.Errorf("UserItem.Gain() = %v, want %v", got, tt.want)
			}
			if u.Quantity != tt.result {
				t.Error("error: quantity not same")
			}
		})
	}
}

func TestUserItem_Consume(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		Item             Item
		Quantity         int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result int
		want   bool
	}{
		{
			name:   "アイテム消費できるか",
			fields: fields{Quantity: 5},
			args:   args{1},
			result: 4,
			want:   true,
		},
		{
			name:   "所持数より大きな値では消費できない",
			fields: fields{Quantity: 5},
			args:   args{6},
			result: 5,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserItem{
				UserResourceBase: tt.fields.UserResourceBase,
				Item:             tt.fields.Item,
				Quantity:         tt.fields.Quantity,
			}
			if got := u.Consume(tt.args.quantity); got != tt.want {
				t.Errorf("UserItem.Consume() = %v, want %v", got, tt.want)
			}
			if u.Quantity != tt.result {
				t.Error("error: quantity not same")
			}
		})
	}
}
