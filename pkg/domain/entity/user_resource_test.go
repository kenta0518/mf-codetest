package entity

import (
	"testing"

	"gorm.io/gorm"
)

func TestUserResourceBase_GetID(t *testing.T) {
	type fields struct {
		Model gorm.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "IDが0の場合",
			fields: fields{
				Model: gorm.Model{ID: 0},
			},
			want: 0,
		},
		{
			name: "IDが0でない場合",
			fields: fields{
				Model: gorm.Model{ID: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserResourceBase{
				Model: tt.fields.Model,
			}
			if got := r.GetID(); got != tt.want {
				t.Errorf("UserResourceBase.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserResourceBase_IsEmpty(t *testing.T) {
	type fields struct {
		Model gorm.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IDが0の場合",
			fields: fields{
				Model: gorm.Model{ID: 0},
			},
			want: true,
		},
		{
			name: "IDが0でない場合",
			fields: fields{
				Model: gorm.Model{ID: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserResourceBase{
				Model: tt.fields.Model,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("UserResourceBase.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
