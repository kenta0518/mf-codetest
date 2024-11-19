package entity

import (
	"reflect"
	"testing"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestNewSenderGiftGachaLog(t *testing.T) {
	type args struct {
		mirrativID string
		giftItem   GiftItem
		txLog      GiftTransactionLog
		results    []RewardContent
	}
	tests := []struct {
		name string
		args args
		want *GiftGachaLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				mirrativID: "user",
				giftItem:   GiftItem{SeedBase: SeedBase{ID: 1001}},
				txLog:      GiftTransactionLog{LogBase: LogBase{gorm.Model{ID: 1}}},
				results:    []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
			},
			want: &GiftGachaLog{
				MirrativID:           "user",
				GiftItemID:           1001,
				GiftItem:             GiftItem{SeedBase: SeedBase{ID: 1001}},
				GiftTransactionLogID: 1,
				GiftTransactionLog:   GiftTransactionLog{LogBase: LogBase{gorm.Model{ID: 1}}},
				IsSender:             true,
				Result:               datatypes.NewJSONType(GiftGachaResultLog{Results: []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}}}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSenderGiftGachaLog(tt.args.mirrativID, tt.args.giftItem, tt.args.txLog, tt.args.results); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSenderGiftGachaLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStreamerGiftGachaLog(t *testing.T) {
	type args struct {
		mirrativID string
		giftItem   GiftItem
		txLog      GiftTransactionLog
		results    []RewardContent
	}
	tests := []struct {
		name string
		args args
		want *GiftGachaLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				mirrativID: "user",
				giftItem:   GiftItem{SeedBase: SeedBase{ID: 1001}},
				txLog:      GiftTransactionLog{LogBase: LogBase{gorm.Model{ID: 1}}},
				results:    []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
			},
			want: &GiftGachaLog{
				MirrativID:           "user",
				GiftItemID:           1001,
				GiftItem:             GiftItem{SeedBase: SeedBase{ID: 1001}},
				GiftTransactionLogID: 1,
				GiftTransactionLog:   GiftTransactionLog{LogBase: LogBase{gorm.Model{ID: 1}}},
				IsSender:             false,
				Result:               datatypes.NewJSONType(GiftGachaResultLog{Results: []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}}}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStreamerGiftGachaLog(tt.args.mirrativID, tt.args.giftItem, tt.args.txLog, tt.args.results); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStreamerGiftGachaLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
