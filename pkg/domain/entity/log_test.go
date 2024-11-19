package entity

import (
	"testing"

	"gorm.io/gorm"
)

func TestLogBase_GetID(t *testing.T) {
	type fields struct {
		Model gorm.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "IDが返る",
			fields: fields{
				Model: gorm.Model{ID: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogBase{
				Model: tt.fields.Model,
			}
			if got := l.GetID(); got != tt.want {
				t.Errorf("LogBase.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogBase_IsEmpty(t *testing.T) {
	type fields struct {
		Model gorm.Model
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IDが0ならtrue",
			fields: fields{
				Model: gorm.Model{ID: 0},
			},
			want: true,
		},
		{
			name: "IDが0じゃないならfalse",
			fields: fields{
				Model: gorm.Model{ID: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogBase{
				Model: tt.fields.Model,
			}
			if got := l.IsEmpty(); got != tt.want {
				t.Errorf("LogBase.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
