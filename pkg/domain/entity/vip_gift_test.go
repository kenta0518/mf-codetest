package entity

import (
	"reflect"
	"testing"
)

func TestVipGift_RewardContents(t *testing.T) {
	type fields struct {
		SeedBase         SeedBase
		VipID            uint
		ContentType1     ContentType
		ContentId1       uint
		ContentQuantity1 int
		ContentType2     ContentType
		ContentId2       uint
		ContentQuantity2 int
		ContentType3     ContentType
		ContentId3       uint
		ContentQuantity3 int
	}
	tests := []struct {
		name   string
		fields fields
		want   []RewardContent
	}{
		{
			name: "有効なRewardContent返せるか",
			fields: fields{
				ContentType1: CONTENT_TYPE_GOLD, ContentId1: 0, ContentQuantity1: 1000,
				ContentType3: CONTENT_TYPE_ITEM, ContentId3: 1001, ContentQuantity3: 5,
			},
			want: []RewardContent{
				{CONTENT_TYPE_GOLD, 0, 1000},
				{CONTENT_TYPE_ITEM, 1001, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := VipGift{
				SeedBase:         tt.fields.SeedBase,
				VipID:            tt.fields.VipID,
				ContentType1:     tt.fields.ContentType1,
				ContentId1:       tt.fields.ContentId1,
				ContentQuantity1: tt.fields.ContentQuantity1,
				ContentType2:     tt.fields.ContentType2,
				ContentId2:       tt.fields.ContentId2,
				ContentQuantity2: tt.fields.ContentQuantity2,
				ContentType3:     tt.fields.ContentType3,
				ContentId3:       tt.fields.ContentId3,
				ContentQuantity3: tt.fields.ContentQuantity3,
			}
			if got := v.RewardContents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VipGift.RewardContents() = %v, want %v", got, tt.want)
			}
		})
	}
}
