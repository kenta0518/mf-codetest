package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserMission(t *testing.T) {
	type args struct {
		userID    uint
		missionID uint
		now       time.Time
	}
	tests := []struct {
		name string
		args args
		want *UserMission
	}{
		{
			name: "UserMissionを生成する",
			args: args{
				userID:    1,
				missionID: 1,
				now:       time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: &UserMission{
				UserID:     1,
				MissionID:  1,
				Progress:   0,
				ReceivedAt: nil,
				ResetAt:    &DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserMission(tt.args.userID, tt.args.missionID, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserMission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMission_IsReceived(t *testing.T) {
	type fields struct {
		ReceivedAt *DateTime
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "受け取っていない",
			fields: fields{
				ReceivedAt: nil,
			},
			want: false,
		},
		{
			name: "受け取っている",
			fields: fields{
				ReceivedAt: &DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserMission{
				ReceivedAt: tt.fields.ReceivedAt,
			}
			if got := u.IsReceived(); got != tt.want {
				t.Errorf("UserMission.IsReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMission_Receive(t *testing.T) {
	type fields struct {
		ReceivedAt *DateTime
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
				ReceivedAt: nil,
			},
			args: args{
				time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserMission{
				ReceivedAt: tt.fields.ReceivedAt,
			}
			u.Receive(tt.args.time)

			if u.ReceivedAt == nil {
				t.Errorf("Receive() = %v, want %v", u.ReceivedAt, tt.args.time)
			}
		})
	}
}

func TestUserMission_UpdateProgress(t *testing.T) {
	type fields struct {
		Progress int
	}
	type args struct {
		progress int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "進捗を更新する",
			fields: fields{
				Progress: 0,
			},
			args: args{
				progress: 1,
			},
			want: true,
		},
		{
			name: "進捗を更新しない",
			fields: fields{
				Progress: 1,
			},
			args: args{
				progress: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserMission{
				Progress: tt.fields.Progress,
			}
			if got := u.UpdateProgress(tt.args.progress); got != tt.want {
				t.Errorf("UserMission.UpdateProgress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMission_AddProgress(t *testing.T) {
	type fields struct {
		Progress int
	}
	type args struct {
		progress int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "進捗を追加する",
			fields: fields{
				Progress: 0,
			},
			args: args{
				progress: 1,
			},
		},
		{
			name: "進捗を追加しない1",
			fields: fields{
				Progress: 1,
			},
			args: args{
				progress: 0,
			},
		},
		{
			name: "進捗を追加しない2",
			fields: fields{
				Progress: 1,
			},
			args: args{
				progress: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserMission{
				Progress: tt.fields.Progress,
			}
			u.AddProgress(tt.args.progress)

			if u.Progress != 1 {
				t.Errorf("AddProgress() = %v, want %v", u.Progress, tt.fields.Progress+tt.args.progress)
			}
		})
	}
}

func TestUserMission_IsAchieved(t *testing.T) {
	type fields struct {
		Mission  Mission
		Progress int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "達成していない",
			fields: fields{
				Mission:  Mission{ConditionValue1: 1},
				Progress: 0,
			},
			want: false,
		},
		{
			name: "達成している",
			fields: fields{
				Mission:  Mission{ConditionValue1: 1},
				Progress: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserMission{
				Mission:  tt.fields.Mission,
				Progress: tt.fields.Progress,
			}
			if got := u.IsAchieved(); got != tt.want {
				t.Errorf("UserMission.IsAchieved() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMission_IsInPeriod(t *testing.T) {
	type fields struct {
		Mission Mission
		ResetAt *DateTime
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "デイリー期間内",
			fields: fields{
				Mission: Mission{MissionGroup: MissionGroup{GroupKind: MissionGroupKindDaily}},
				ResetAt: &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "デイリー期間外1",
			fields: fields{
				Mission: Mission{MissionGroup: MissionGroup{GroupKind: MissionGroupKindDaily}},
				ResetAt: &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 2, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "デイリー期間外2",
			fields: fields{
				Mission: Mission{MissionGroup: MissionGroup{GroupKind: MissionGroupKindDaily}},
				ResetAt: &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 4, 30, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "デイリー以外",
			fields: fields{
				Mission: Mission{MissionGroup: MissionGroup{GroupKind: MissionGroupKindMain}},
				ResetAt: &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 2, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserMission{
				Mission: tt.fields.Mission,
				ResetAt: tt.fields.ResetAt,
			}
			if got := u.IsInPeriod(tt.args.t); got != tt.want {
				t.Errorf("UserMission.IsInPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}
