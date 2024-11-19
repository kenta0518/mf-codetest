package entity

import (
	"reflect"
	"testing"
)

func TestNewUserMainQuestGroup(t *testing.T) {
	type args struct {
		userId       uint
		questGroupId uint
		unlock       bool
	}
	tests := []struct {
		name string
		args args
		want *UserMainQuestGroup
	}{
		{
			name: "UserQuestGroup作成",
			args: args{
				userId:       1,
				questGroupId: 1,
				unlock:       false,
			},
			want: &UserMainQuestGroup{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock: false,
			},
		},
		{
			name: "作成時に解放",
			args: args{
				userId:       1,
				questGroupId: 1,
				unlock:       true,
			},
			want: &UserMainQuestGroup{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserQuestGroup(tt.args.userId, tt.args.questGroupId, tt.args.unlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserQuestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMainQuestGroup_Unlock(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		QuestGroup       MainQuestGroup
		UnLock           bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "解放",
			fields: fields{
				UnLock: false,
			},
		},
		{
			name: "解放済みは変化なし",
			fields: fields{
				UnLock: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserMainQuestGroup{
				UserResourceBase: tt.fields.UserResourceBase,
				MainQuestGroup:   tt.fields.QuestGroup,
				UnLock:           tt.fields.UnLock,
			}
			q.Unlock()
		})
	}
}
