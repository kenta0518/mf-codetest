package entity

import (
	"reflect"
	"testing"
)

func TestNewUserMainQuest(t *testing.T) {
	type args struct {
		userId  uint
		questId uint
		unlock  bool
	}
	tests := []struct {
		name string
		args args
		want *UserMainQuest
	}{
		{
			name: "ユーザークエスト作成",
			args: args{
				userId:  1,
				questId: 1,
				unlock:  false,
			},
			want: &UserMainQuest{
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
				userId:  1,
				questId: 1,
				unlock:  true,
			},
			want: &UserMainQuest{
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
			if got := NewUserMainQuest(tt.args.userId, tt.args.questId, tt.args.unlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserQuest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMainQuest_Unlock(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		Quest            MainQuest
		UnLock           bool
		IsClear          bool
		HighScore        uint
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
			q := &UserMainQuest{
				UserResourceBase: tt.fields.UserResourceBase,
				MainQuest:        tt.fields.Quest,
				UnLock:           tt.fields.UnLock,
				IsClear:          tt.fields.IsClear,
				HighScore:        tt.fields.HighScore,
			}
			q.Unlock()

			if !q.UnLock {
				t.Errorf("Unlock() = %v, want %v", q.UnLock, true)
			}
		})
	}
}

func TestUserMainQuest_Clear(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		Quest            MainQuest
		UnLock           bool
		IsClear          bool
		HighScore        uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "クリア",
			fields: fields{
				IsClear: false,
			},
		},
		{
			name: "クリア済みは変化なし",
			fields: fields{
				IsClear: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserMainQuest{
				UserResourceBase: tt.fields.UserResourceBase,
				MainQuest:        tt.fields.Quest,
				UnLock:           tt.fields.UnLock,
				IsClear:          tt.fields.IsClear,
				HighScore:        tt.fields.HighScore,
			}
			q.Clear()

			if !q.IsClear {
				t.Errorf("Clear() = %v, want %v", q.IsClear, true)
			}
		})
	}
}

func TestUserMainQuest_SetHighScore(t *testing.T) {
	var testScore uint = 1000
	type fields struct {
		UserResourceBase UserResourceBase
		Quest            MainQuest
		UnLock           bool
		IsClear          bool
		HighScore        uint
	}
	type args struct {
		score uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ハイスコア未設定",
			fields: fields{
				HighScore: 0,
			},
			args: args{
				score: testScore,
			},
		},
		{
			name: "ハイスコア未更新",
			fields: fields{
				HighScore: testScore,
			},
			args: args{
				score: 500,
			},
		},
		{
			name: "ハイスコア更新",
			fields: fields{
				HighScore: 500,
			},
			args: args{
				score: testScore,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserMainQuest{
				UserResourceBase: tt.fields.UserResourceBase,
				MainQuest:        tt.fields.Quest,
				UnLock:           tt.fields.UnLock,
				IsClear:          tt.fields.IsClear,
				HighScore:        tt.fields.HighScore,
			}
			q.SetHighScore(tt.args.score)

			if q.HighScore != testScore {
				t.Errorf("SetHighScore() = %v, want %v", q.HighScore, tt.args.score)
			}
		})
	}
}

func TestUserMainQuest_SetClearRank(t *testing.T) {
	highRank := QuestScoreRankS
	type fields struct {
		UserResourceBase UserResourceBase
		MainQuest        MainQuest
		MainQuestGroupID uint
		UnLock           bool
		IsClear          bool
		HighScore        uint
		ClearRank        QuestScoreRankType
	}
	type args struct {
		rank QuestScoreRankType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "クリアランク未設定",
			fields: fields{
				ClearRank: QuestScoreRankNone,
			},
			args: args{
				rank: highRank,
			},
		},
		{
			name: "クリアランク未更新",
			fields: fields{
				ClearRank: highRank,
			},
			args: args{
				rank: QuestScoreRankA,
			},
		},
		{
			name: "クリアランク更新",
			fields: fields{
				ClearRank: QuestScoreRankA,
			},
			args: args{
				rank: highRank,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserMainQuest{
				UserResourceBase: tt.fields.UserResourceBase,
				MainQuest:        tt.fields.MainQuest,
				UnLock:           tt.fields.UnLock,
				IsClear:          tt.fields.IsClear,
				HighScore:        tt.fields.HighScore,
				ClearRank:        tt.fields.ClearRank,
			}
			q.SetClearRank(tt.args.rank)

			if q.ClearRank != highRank {
				t.Errorf("SetClearRank() = %v, want %v", q.ClearRank, highRank)
			}
		})
	}
}
