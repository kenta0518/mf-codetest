package entity

import (
	"reflect"
	"testing"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestNewBonusGiftGachaStockLog(t *testing.T) {
	type args struct {
		txLog     *GiftTransactionLog
		giftItem  *GiftItem
		bannerID  uint
		degree1ID uint
		degree2ID uint
		degree3ID uint
		result    []RewardContent
	}
	tests := []struct {
		name string
		args args
		want *BonusGiftGachaStockLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				txLog: &GiftTransactionLog{
					LogBase:  LogBase{gorm.Model{ID: 1}},
					SenderID: "sender_id",
				},
				giftItem: &GiftItem{
					SeedBase: SeedBase{ID: 1000},
				},
				bannerID:  1,
				degree1ID: 2,
				degree2ID: 3,
				degree3ID: 4,
				result: []RewardContent{
					{ContentType: CONTENT_TYPE_ITEM, ContentID: 1, ContentQuantity: 1},
				},
			},
			want: &BonusGiftGachaStockLog{
				MirrativID:           "sender_id",
				GiftTransactionLogID: 1,
				GiftTransactionLog:   GiftTransactionLog{LogBase: LogBase{gorm.Model{ID: 1}}, SenderID: "sender_id"},
				GiftItemID:           1000,
				GiftItem:             GiftItem{SeedBase: SeedBase{ID: 1000}},
				ProfileBannerID:      1,
				ProfileDegree1ID:     2,
				ProfileDegree2ID:     3,
				ProfileDegree3ID:     4,
				Stock:                1,
				Result: datatypes.NewJSONType(BonusGiftGachaResultLog{
					Results: []RewardContent{
						{ContentType: CONTENT_TYPE_ITEM, ContentID: 1, ContentQuantity: 1},
					},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBonusGiftGachaStockLog(tt.args.txLog, tt.args.giftItem, tt.args.bannerID, tt.args.degree1ID, tt.args.degree2ID, tt.args.degree3ID, tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBonusGiftGachaStockLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
