package entity

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
)

func TestTerm_IsInTerm(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	type fields struct {
		StartAt     DateTime
		EndAt       DateTime
		TestStartAt DateTime
		TestEndAt   DateTime
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
			name:   "期間(通常ユーザ)",
			fields: fields{StartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)}, EndAt: DateTime{time.Date(2100, 1, 1, 0, 0, 0, 0, loc)}, TestStartAt: DateTime{time.Date(2050, 1, 1, 0, 0, 0, 0, loc)}, TestEndAt: DateTime{time.Date(2150, 1, 1, 0, 0, 0, 0, loc)}},
			args:   args{user: &User{UserKind: Player}},
			want:   true,
		},
		{
			name:   "期間(スーパーユーザ)",
			fields: fields{StartAt: DateTime{time.Date(2050, 1, 1, 0, 0, 0, 0, loc)}, EndAt: DateTime{time.Date(2150, 1, 1, 0, 0, 0, 0, loc)}, TestStartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)}, TestEndAt: DateTime{time.Date(2100, 1, 1, 0, 0, 0, 0, loc)}},
			args:   args{user: &User{UserKind: SuperUser}},
			want:   true,
		},
		{
			name:   "時差あり",
			fields: fields{StartAt: DateTime{time.Date(2023, 1, 1, 0, 0, 0, 0, loc)}, EndAt: DateTime{time.Date(2023, 1, 2, 0, 0, 0, 0, loc)}, TestStartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)}, TestEndAt: DateTime{time.Date(2100, 1, 1, 0, 0, 0, 0, loc)}},
			args:   args{user: &User{TimeDifference: 24 * time.Hour, UserKind: Player}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flextime.Fix(time.Date(2023, 1, 1, 12, 0, 0, 0, loc))
			tr := &Term{
				StartAt:     tt.fields.StartAt,
				EndAt:       tt.fields.EndAt,
				TestStartAt: tt.fields.TestStartAt,
				TestEndAt:   tt.fields.TestEndAt,
			}
			if got := tr.IsInTerm(tt.args.user); got != tt.want {
				t.Errorf("Term.IsInTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}
