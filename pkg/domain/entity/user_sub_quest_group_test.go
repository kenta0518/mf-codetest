package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserSubQuestGroup(t *testing.T) {
	type args struct {
		userID          uint
		subQuestGroupID uint
		unlock          bool
		now             time.Time
	}
	tests := []struct {
		name string
		args args
		want *UserSubQuestGroup
	}{
		{
			name: "UserSubQuestGroupを生成する",
			args: args{
				userID:          1,
				subQuestGroupID: 1,
				unlock:          false,
				now:             time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: &UserSubQuestGroup{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock:    false,
				PlayCount: 0,
				ResetAt:   &DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
		},
		{
			name: "開放状態でUserSubQuestGroupを生成する",
			args: args{
				userID:          1,
				subQuestGroupID: 1,
				unlock:          true,
				now:             time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: &UserSubQuestGroup{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock:    true,
				PlayCount: 0,
				ResetAt:   &DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserSubQuestGroup(tt.args.userID, tt.args.subQuestGroupID, tt.args.unlock, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserSubQuestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSubQuestGroup_Unlock(t *testing.T) {
	type fields struct {
		UnLock bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "解放する",
			fields: fields{
				UnLock: false,
			},
		},
		{
			name: "解放済み",
			fields: fields{
				UnLock: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuestGroup{
				UnLock: tt.fields.UnLock,
			}
			u.Unlock()

			if !u.UnLock {
				t.Errorf("Unlock() = %v, want %v", u.UnLock, true)
			}
		})
	}
}

func TestUserSubQuestGroup_Reset(t *testing.T) {
	type fields struct {
		PlayCount uint
		ResetAt   *DateTime
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "リセットする",
			fields: fields{
				PlayCount: 1,
			},
			args: args{
				now: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuestGroup{
				PlayCount: tt.fields.PlayCount,
				ResetAt:   tt.fields.ResetAt,
			}
			u.Reset(tt.args.now)

			if u.PlayCount != 0 {
				t.Errorf("Reset() = %v, want %v", u.PlayCount, 0)
			}
		})
	}
}

func TestUserSubQuestGroup_AddPlayCount(t *testing.T) {
	type fields struct {
		PlayCount uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "プレイ回数を追加する",
			fields: fields{
				PlayCount: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuestGroup{
				PlayCount: tt.fields.PlayCount,
			}
			u.AddPlayCount()

			if u.PlayCount != tt.fields.PlayCount+1 {
				t.Errorf("AddPlayCount() = %v, want %v", u.PlayCount, tt.fields.PlayCount+1)
			}
		})
	}
}

func TestUserSubQuestGroup_IsInPeriod(t *testing.T) {
	type fields struct {
		SubQuestGroup SubQuestGroup
		ResetAt       *DateTime
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
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeDaily},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "デイリー期間外1",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeDaily},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 2, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "デイリー期間外2",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeDaily},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 4, 30, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "ウィークリー期間内",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeWeekly},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 7, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "ウィークリー期間外1",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeWeekly},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 8, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "ウィークリー期間外2",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeWeekly},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 4, 30, 12, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "期間なし",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeNone},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "Term期間中",
			fields: fields{
				SubQuestGroup: SubQuestGroup{PlayLimitType: PlayLimitTypeInTerm},
				ResetAt:       &DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				t: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuestGroup{
				SubQuestGroup: tt.fields.SubQuestGroup,
				ResetAt:       tt.fields.ResetAt,
			}
			if got := u.IsInPeriod(tt.args.t); got != tt.want {
				t.Errorf("UserSubQuestGroup.IsInPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}
