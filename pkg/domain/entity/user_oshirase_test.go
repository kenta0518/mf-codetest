package entity

import (
	"testing"

	"gorm.io/gorm"
)

func TestUserOshirase_Read(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		UserID     uint
		OshiraseID uint
		Oshirase   Oshirase
		IsRead     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *UserOshirase
	}{
		{
			name: "Read",
			fields: fields{
				Model:      gorm.Model{},
				UserID:     1,
				OshiraseID: 1,
				Oshirase:   Oshirase{},
				IsRead:     false,
			},
			want: &UserOshirase{
				Model:      gorm.Model{},
				UserID:     1,
				OshiraseID: 1,
				Oshirase:   Oshirase{},
				IsRead:     true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &UserOshirase{
				Model:      tt.fields.Model,
				UserID:     tt.fields.UserID,
				OshiraseID: tt.fields.OshiraseID,
				Oshirase:   tt.fields.Oshirase,
				IsRead:     tt.fields.IsRead,
			}
			o.Read()
			if o.IsRead != tt.want.IsRead {
				t.Errorf("UserOshirase.Read() = %v, want %v", o.IsRead, tt.want.IsRead)
			}
		})
	}
}
