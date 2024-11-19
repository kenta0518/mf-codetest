package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserEventSeason(t *testing.T) {
	type args struct {
		userID        uint
		eventSeasonID uint
		now           time.Time
	}
	tests := []struct {
		name string
		args args
		want *UserEventSeason
	}{
		{
			name: "UserEventSeasonを生成できる",
			args: args{
				userID:        1,
				eventSeasonID: 1,
				now:           time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
			want: &UserEventSeason{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				EventPoint: 0,
				Stamina:    10,
				CheckedAt:  DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserEventSeason(tt.args.userID, tt.args.eventSeasonID, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserEventSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserEventSeason_AddEventPoint(t *testing.T) {
	type fields struct {
		EventPoint uint
	}
	type args struct {
		point uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "EventPointにポイントを追加できる",
			fields: fields{
				EventPoint: 0,
			},
			args: args{
				point: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserEventSeason{
				EventPoint: tt.fields.EventPoint,
			}
			u.AddEventPoint(tt.args.point)

			if u.EventPoint != tt.args.point+tt.fields.EventPoint {
				t.Errorf("AddEventPoint() = %v, want %v", u.EventPoint, tt.args.point+tt.fields.EventPoint)
			}
		})
	}
}

func TestUserEventSeason_UseStamina(t *testing.T) {
	now := time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)
	type fields struct {
		Stamina   uint
		CheckedAt DateTime
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "スタミナを消費できる",
			fields: fields{
				Stamina: 9,
			},
			args: args{
				now: now,
			},
			want: true,
		},
		{
			name: "スタミナが足りない場合",
			fields: fields{
				Stamina: 0,
			},
			args: args{
				now: now,
			},
			want: false,
		},
		{
			name: "スタミナが最大",
			fields: fields{
				Stamina: 10,
			},
			args: args{
				now: now,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserEventSeason{
				Stamina:   tt.fields.Stamina,
				CheckedAt: tt.fields.CheckedAt,
			}
			if got := u.UseStamina(tt.args.now); got != tt.want {
				t.Errorf("UserEventSeason.UseStamina() = %v, want %v", got, tt.want)
			}

			if tt.fields.Stamina == 10 && u.CheckedAt.Time != now {
				t.Errorf("UserEventSeason.UseStamina() = %v, want %v", u.CheckedAt.Time, now)
			}
		})
	}
}

func TestUserEventSeason_RecoveryStamina(t *testing.T) {
	type fields struct {
		Stamina   uint
		CheckedAt DateTime
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want2  uint
		want3  time.Time
	}{
		{
			name: "Staminaが最大値の場合",
			fields: fields{
				Stamina:   EventSeasonStaminaMax,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want:  false,
			want2: EventSeasonStaminaMax,
			want3: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから10分経過していない",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 5, 0, 0, time.Local),
			},
			want:  false,
			want2: 5,
			want3: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから10分経過している",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 12, 0, 0, time.Local),
			},
			want:  true,
			want2: 6,
			want3: time.Date(2023, 5, 1, 0, 10, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから20分経過している",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 28, 0, 0, time.Local),
			},
			want:  true,
			want2: 7,
			want3: time.Date(2023, 5, 1, 0, 20, 0, 0, time.Local),
		},
		{
			name: "Stamina最大値以上に回復しようとした場合",
			fields: fields{
				Stamina:   0,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want:  true,
			want2: EventSeasonStaminaMax,
			want3: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserEventSeason{
				Stamina:   tt.fields.Stamina,
				CheckedAt: tt.fields.CheckedAt,
			}
			if got := u.RecoveryStamina(tt.args.now); got != tt.want {
				t.Errorf("UserEventSeason.RecoveryStamina() = %v, want %v", got, tt.want)
			}
			if u.Stamina != tt.want2 {
				t.Errorf("UserSoloRaid.ReceveryStamina() = %v, want %v", u.Stamina, tt.want2)
			}
			if u.CheckedAt.Time != tt.want3 {
				t.Errorf("UserSoloRaid.ReceveryStamina() = %v, want %v", u.CheckedAt.Time, tt.want3)
			}
		})
	}
}
