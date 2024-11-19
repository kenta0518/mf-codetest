package entity

import (
	"reflect"
	"testing"
)

func TestGetClearRewardType(t *testing.T) {
	type args struct {
		oldRank QuestScoreRankType
		newRank QuestScoreRankType
	}
	tests := []struct {
		name string
		args args
		want []ClearRewardType
	}{
		{
			name: "RankNone -> RankNone",
			args: args{
				oldRank: QuestScoreRankNone,
				newRank: QuestScoreRankNone,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankNone -> RankC",
			args: args{
				oldRank: QuestScoreRankNone,
				newRank: QuestScoreRankC,
			},
			want: []ClearRewardType{ClearRewardRankC},
		},
		{
			name: "RankNone -> RankB",
			args: args{
				oldRank: QuestScoreRankNone,
				newRank: QuestScoreRankB,
			},
			want: []ClearRewardType{ClearRewardRankC, ClearRewardRankB},
		},
		{
			name: "RankNone -> RankA",
			args: args{
				oldRank: QuestScoreRankNone,
				newRank: QuestScoreRankA,
			},
			want: []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA},
		},
		{
			name: "RankNone -> RankS",
			args: args{
				oldRank: QuestScoreRankNone,
				newRank: QuestScoreRankS,
			},
			want: []ClearRewardType{ClearRewardRankC, ClearRewardRankB, ClearRewardRankA, ClearRewardRankS},
		},
		{
			name: "RankC -> RankNone",
			args: args{
				oldRank: QuestScoreRankC,
				newRank: QuestScoreRankNone,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankC -> RankC",
			args: args{
				oldRank: QuestScoreRankC,
				newRank: QuestScoreRankC,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankC -> RankB",
			args: args{
				oldRank: QuestScoreRankC,
				newRank: QuestScoreRankB,
			},
			want: []ClearRewardType{ClearRewardRankB},
		},
		{
			name: "RankC -> RankA",
			args: args{
				oldRank: QuestScoreRankC,
				newRank: QuestScoreRankA,
			},
			want: []ClearRewardType{ClearRewardRankB, ClearRewardRankA},
		},
		{
			name: "RankC -> RankS",
			args: args{
				oldRank: QuestScoreRankC,
				newRank: QuestScoreRankS,
			},
			want: []ClearRewardType{ClearRewardRankB, ClearRewardRankA, ClearRewardRankS},
		},
		{
			name: "RankB -> RankNone",
			args: args{
				oldRank: QuestScoreRankB,
				newRank: QuestScoreRankNone,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankB -> RankC",
			args: args{
				oldRank: QuestScoreRankB,
				newRank: QuestScoreRankC,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankB -> RankB",
			args: args{
				oldRank: QuestScoreRankB,
				newRank: QuestScoreRankB,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankB -> RankA",
			args: args{
				oldRank: QuestScoreRankB,
				newRank: QuestScoreRankA,
			},
			want: []ClearRewardType{ClearRewardRankA},
		},
		{
			name: "RankB -> RankS",
			args: args{
				oldRank: QuestScoreRankB,
				newRank: QuestScoreRankS,
			},
			want: []ClearRewardType{ClearRewardRankA, ClearRewardRankS},
		},
		{
			name: "RankA -> RankNone",
			args: args{
				oldRank: QuestScoreRankA,
				newRank: QuestScoreRankNone,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankA -> RankC",
			args: args{
				oldRank: QuestScoreRankA,
				newRank: QuestScoreRankC,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankA -> RankB",
			args: args{
				oldRank: QuestScoreRankA,
				newRank: QuestScoreRankB,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankA -> RankA",
			args: args{
				oldRank: QuestScoreRankA,
				newRank: QuestScoreRankA,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankA -> RankS",
			args: args{
				oldRank: QuestScoreRankA,
				newRank: QuestScoreRankS,
			},
			want: []ClearRewardType{ClearRewardRankS},
		},
		{
			name: "RankS -> RankNone",
			args: args{
				oldRank: QuestScoreRankS,
				newRank: QuestScoreRankNone,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankS -> RankC",
			args: args{
				oldRank: QuestScoreRankS,
				newRank: QuestScoreRankC,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankS -> RankB",
			args: args{
				oldRank: QuestScoreRankS,
				newRank: QuestScoreRankB,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankS -> RankA",
			args: args{
				oldRank: QuestScoreRankS,
				newRank: QuestScoreRankA,
			},
			want: []ClearRewardType{},
		},
		{
			name: "RankS -> RankS",
			args: args{
				oldRank: QuestScoreRankS,
				newRank: QuestScoreRankS,
			},
			want: []ClearRewardType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClearRewardType(tt.args.oldRank, tt.args.newRank); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClearRewardType() = %v, want %v", got, tt.want)
			}
		})
	}
}
