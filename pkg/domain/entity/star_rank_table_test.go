package entity

import (
	"reflect"
	"testing"
)

func TestStarRankTable_GetRewardContent(t *testing.T) {
	type fields struct {
		SeedBase        SeedBase
		Rank            uint
		RankExp         uint
		ContentType     ContentType
		ContentId       uint
		ContentQuantity int
	}
	tests := []struct {
		name   string
		fields fields
		want   RewardContent
	}{
		{
			name:   "RewardContentの取得ができる",
			fields: fields{ContentType: CONTENT_TYPE_ITEM, ContentId: 1, ContentQuantity: 1},
			want:   RewardContent{ContentType: CONTENT_TYPE_ITEM, ContentID: 1, ContentQuantity: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StarRankTable{
				SeedBase:        tt.fields.SeedBase,
				Rank:            tt.fields.Rank,
				RankMaxExp:      tt.fields.RankExp,
				ContentType:     tt.fields.ContentType,
				ContentId:       tt.fields.ContentId,
				ContentQuantity: tt.fields.ContentQuantity,
			}
			if got := s.GetRewardContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StarRankTable.GetRewardContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
