package entity

import (
	"testing"

	"gorm.io/gorm"
)

func TestUserGiftProperty_AddGiftEventStamina(t *testing.T) {
	type fields struct {
		Model            gorm.Model
		MirrativID       string
		GiftEventStamina uint
	}
	type args struct {
		stamina uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserGiftProperty
	}{
		{
			name: "イベントスタミナが加算される",
			fields: fields{
				GiftEventStamina: 100,
			},
			args: args{
				stamina: 100,
			},
			want: &UserGiftProperty{
				GiftEventStamina: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserGiftProperty{
				Model:            tt.fields.Model,
				MirrativID:       tt.fields.MirrativID,
				GiftEventStamina: tt.fields.GiftEventStamina,
			}
			u.AddGiftEventStamina(tt.args.stamina)
		})
	}
}
