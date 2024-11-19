package entity

import "testing"

func TestQuestScoreRankTypeIndex(t *testing.T) {
	type args struct {
		rank QuestScoreRankType
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ランクなし",
			args: args{
				rank: QuestScoreRankNone,
			},
			want: 0,
		},
		{
			name: "ランクC",
			args: args{
				rank: QuestScoreRankC,
			},
			want: 1,
		},
		{
			name: "ランクB",
			args: args{
				rank: QuestScoreRankB,
			},
			want: 2,
		},
		{
			name: "ランクA",
			args: args{
				rank: QuestScoreRankA,
			},
			want: 3,
		},
		{
			name: "ランクS",
			args: args{
				rank: QuestScoreRankS,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuestScoreRankTypeIndex(tt.args.rank); got != tt.want {
				t.Errorf("QuestScoreRankTypeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
