package entity

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

func TestSubQuestGroup_IsOpen(t *testing.T) {
	flextime.Fix(time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local))
	type fields struct {
		OpenWeek int
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "常にオープン",
			fields: fields{
				OpenWeek: 0b1111111,
			},
			args: args{
				user: nil,
			},
			want: true,
		},
		{
			name: "火曜日にオープン",
			fields: fields{
				OpenWeek: 0b0000010,
			},
			args: args{
				user: nil,
			},
			want: true,
		},
		{
			name: "火曜日にオープンしない",
			fields: fields{
				OpenWeek: 0b0000001,
			},
			args: args{
				user: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubQuestGroup{
				OpenWeek: tt.fields.OpenWeek,
			}
			if got := s.IsOpen(tt.args.user); got != tt.want {
				t.Errorf("SubQuestGroup.IsOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}
