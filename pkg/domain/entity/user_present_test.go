package entity

import (
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"gorm.io/gorm"
)

func TestUserPresent_CanBeReceived(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	flextime.Fix(time.Date(2023, 3, 1, 0, 0, 0, 0, loc))
	type fields struct {
		Model        gorm.Model
		Term         Term
		PresentID    *uint
		Present      *Present
		Condition    string
		ReceivedTime *time.Time
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
			name: "受け取り可能(通常ユーザ)",
			fields: fields{
				Term: Term{
					StartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)},
					EndAt:   DateTime{time.Date(3000, 1, 1, 0, 0, 0, 0, loc)},
				},
			},
			args: args{
				user: &User{UserKind: Player},
			},
			want: true,
		},
		{
			name: "期間外(通常ユーザ)",
			fields: fields{
				Term: Term{
					StartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)},
					EndAt:   DateTime{time.Date(2010, 1, 1, 0, 0, 0, 0, loc)},
				},
			},
			args: args{
				user: &User{UserKind: Player},
			},
			want: false,
		},
		{
			name: "受け取り可能(スーパーユーザ)",
			fields: fields{
				Term: Term{
					TestStartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)},
					TestEndAt:   DateTime{time.Date(3000, 1, 1, 0, 0, 0, 0, loc)},
				},
			},
			args: args{
				user: &User{UserKind: SuperUser},
			},
			want: true,
		},
		{
			name: "期間外(スーパーユーザ)",
			fields: fields{
				Term: Term{
					TestStartAt: DateTime{time.Date(2000, 1, 1, 0, 0, 0, 0, loc)},
					TestEndAt:   DateTime{time.Date(2010, 1, 1, 0, 0, 0, 0, loc)},
				},
			},
			args: args{
				user: &User{UserKind: SuperUser},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPresent{
				Model:        tt.fields.Model,
				Term:         tt.fields.Term,
				PresentID:    tt.fields.PresentID,
				Present:      tt.fields.Present,
				Condition:    tt.fields.Condition,
				ReceivedTime: tt.fields.ReceivedTime,
			}
			if got := u.CanBeReceived(tt.args.user); got != tt.want {
				t.Errorf("UserPresent.CanBeReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPresent_IsReceived(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	type fields struct {
		Model        gorm.Model
		Term         Term
		PresentID    *uint
		Present      *Present
		Condition    string
		ReceivedTime *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "受け取り済み",
			fields: fields{
				ReceivedTime: func(t time.Time) *time.Time { return &t }(time.Date(2023, 3, 1, 0, 0, 0, 0, loc)),
			},
			want: true,
		},
		{
			name:   "未受け取り",
			fields: fields{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPresent{
				Model:        tt.fields.Model,
				Term:         tt.fields.Term,
				PresentID:    tt.fields.PresentID,
				Present:      tt.fields.Present,
				Condition:    tt.fields.Condition,
				ReceivedTime: tt.fields.ReceivedTime,
			}
			if got := u.IsReceived(); got != tt.want {
				t.Errorf("UserPresent.IsReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserPresent_BeReceived(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	flextime.Fix(time.Date(2023, 3, 1, 0, 0, 0, 0, loc))
	type fields struct {
		Model        gorm.Model
		Term         Term
		UserID       uint
		User         User
		PresentID    *uint
		Present      *Present
		Condition    string
		ReceivedTime *time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "受け取る",
			fields: fields{
				ReceivedTime: func(t time.Time) *time.Time { return &t }(flextime.Now()),
			},
			args: args{
				time: time.Date(2023, 3, 1, 0, 0, 0, 0, loc),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserPresent{
				Model:        tt.fields.Model,
				Term:         tt.fields.Term,
				PresentID:    tt.fields.PresentID,
				Present:      tt.fields.Present,
				Condition:    tt.fields.Condition,
				ReceivedTime: tt.fields.ReceivedTime,
			}
			u.BeReceived(tt.args.time)
		})
	}
}
