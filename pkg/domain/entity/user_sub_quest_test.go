package entity

import (
	"reflect"
	"testing"
)

func TestNewUserSubQuest(t *testing.T) {
	type args struct {
		userId     uint
		subQuestId uint
		unlock     bool
	}
	tests := []struct {
		name string
		args args
		want *UserSubQuest
	}{
		{
			name: "UserSubQuestを生成する",
			args: args{
				userId:     1,
				subQuestId: 1,
				unlock:     false,
			},
			want: &UserSubQuest{
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
			name: "解放状態でUserSubQuestを生成する",
			args: args{
				userId:     1,
				subQuestId: 1,
				unlock:     true,
			},
			want: &UserSubQuest{
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
			if got := NewUserSubQuest(tt.args.userId, tt.args.subQuestId, tt.args.unlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserSubQuest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSubQuest_Unlock(t *testing.T) {
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
			name: "解放済みの場合は変更しない",
			fields: fields{
				UnLock: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuest{
				UnLock: tt.fields.UnLock,
			}
			u.Unlock()

			if !u.UnLock {
				t.Errorf("Unlock() = %v, want %v", u.UnLock, true)
			}
		})
	}
}

func TestUserSubQuest_Clear(t *testing.T) {
	type fields struct {
		IsClear bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "クリアする",
			fields: fields{
				IsClear: false,
			},
		},
		{
			name: "クリア済みの場合は変更しない",
			fields: fields{
				IsClear: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuest{
				IsClear: tt.fields.IsClear,
			}
			u.Clear()
		})
	}
}

func TestUserSubQuest_SetHighScore(t *testing.T) {
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
			name: "ハイスコアを更新する",
			fields: fields{
				HighScore: 0,
			},
			args: args{
				score: highScore,
			},
		},
		{
			name: "ハイスコアを更新しない",
			fields: fields{
				HighScore: highScore,
			},
			args: args{
				score: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSubQuest{
				HighScore: tt.fields.HighScore,
			}
			u.SetHighScore(tt.args.score)

			if u.HighScore != highScore {
				t.Errorf("SetHighScore() = %v, want %v", u.HighScore, highScore)
			}
		})
	}
}

func TestUserSubQuest_SetClearRank(t *testing.T) {
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
			u := &UserSubQuest{
				ClearRank: tt.fields.ClearRank,
			}
			u.SetClearRank(tt.args.rank)

			if u.ClearRank != highRank {
				t.Errorf("SetClearRank() = %v, want %v", u.ClearRank, highRank)
			}
		})
	}
}
