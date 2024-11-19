package entity

import (
	"reflect"
	"testing"
)

func TestNewUserEventQuest(t *testing.T) {
	type args struct {
		userId  uint
		questId uint
		unlock  bool
	}
	tests := []struct {
		name string
		args args
		want *UserEventQuest
	}{
		{
			name: "UserEventQuestを生成できる",
			args: args{
				userId:  1,
				questId: 1,
				unlock:  false,
			},
			want: &UserEventQuest{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock:    false,
				IsClear:   false,
				HighScore: 0,
				ClearRank: QuestScoreRankNone,
			},
		},
		{
			name: "Unlock状態で生成できる",
			args: args{
				userId:  1,
				questId: 1,
				unlock:  true,
			},
			want: &UserEventQuest{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				UnLock:    true,
				IsClear:   false,
				HighScore: 0,
				ClearRank: QuestScoreRankNone,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserEventQuest(tt.args.userId, tt.args.questId, tt.args.unlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserEventQuest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserEventQuest_Unlock(t *testing.T) {
	type fields struct {
		UnLock bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Unlockできる",
			fields: fields{
				UnLock: false,
			},
		},
		{
			name: "Unlock済みは何もしない",
			fields: fields{
				UnLock: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserEventQuest{
				UnLock: tt.fields.UnLock,
			}
			q.Unlock()

			if !q.UnLock {
				t.Errorf("Unlock() = %v, want %v", q.UnLock, true)
			}
		})
	}
}

func TestUserEventQuest_Clear(t *testing.T) {
	type fields struct {
		IsClear bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "クリアできる",
			fields: fields{
				IsClear: false,
			},
		},
		{
			name: "クリア済みは何もしない",
			fields: fields{
				IsClear: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserEventQuest{
				IsClear: tt.fields.IsClear,
			}
			q.Clear()

			if !q.IsClear {
				t.Errorf("Clear() = %v, want %v", q.IsClear, true)
			}
		})
	}
}

func TestUserEventQuest_SetHighScore(t *testing.T) {
	highScore := uint(1000)
	type fields struct {
		HighScore uint
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
			name: "ハイスコアを更新できる",
			fields: fields{
				HighScore: 0,
			},
			args: args{
				score: highScore,
			},
		},
		{
			name: "ハイスコアの更新がない",
			fields: fields{
				HighScore: highScore,
			},
			args: args{
				score: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &UserEventQuest{
				HighScore: tt.fields.HighScore,
			}
			q.SetHighScore(tt.args.score)

			if q.HighScore != highScore {
				t.Errorf("SetHighScore() = %v, want %v", q.HighScore, highScore)
			}
		})
	}
}

func TestUserEventQuest_SetClearRank(t *testing.T) {
	highRank := QuestScoreRankS
	type fields struct {
		ClearRank QuestScoreRankType
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
			q := &UserEventQuest{
				ClearRank: tt.fields.ClearRank,
			}
			q.SetClearRank(tt.args.rank)

			if q.ClearRank != highRank {
				t.Errorf("SetClearRank() = %v, want %v", q.ClearRank, highRank)
			}
		})
	}
}
