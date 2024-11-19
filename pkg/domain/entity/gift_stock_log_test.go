package entity

import (
	"reflect"
	"testing"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestNewGiftStockLog(t *testing.T) {
	type args struct {
		txLog    *GiftTransactionLog
		giftItem *GiftItem
	}
	tests := []struct {
		name string
		args args
		want *GiftStockLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				txLog: &GiftTransactionLog{
					LogBase:    LogBase{gorm.Model{ID: 10}},
					StreamerID: "user1",
					ItemNum:    2,
				},
				giftItem: &GiftItem{
					SeedBase: SeedBase{ID: 10},
				},
			},
			want: &GiftStockLog{
				MirrativID:           "user1",
				GiftTransactionLogID: 10,
				GiftTransactionLog: GiftTransactionLog{
					LogBase:    LogBase{gorm.Model{ID: 10}},
					StreamerID: "user1",
					ItemNum:    2,
				},
				GiftItemID: 10,
				GiftItem: GiftItem{
					SeedBase: SeedBase{ID: 10},
				},
				Stock: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGiftStockLog(tt.args.txLog, tt.args.giftItem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGiftStockLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGachaGiftStockLog(t *testing.T) {
	type args struct {
		txLog          *GiftTransactionLog
		giftItem       *GiftItem
		streamerResult []RewardContent
		senderResult   []RewardContent
	}
	tests := []struct {
		name string
		args args
		want *GiftStockLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				txLog: &GiftTransactionLog{
					LogBase:    LogBase{gorm.Model{ID: 10}},
					StreamerID: "user1",
					ItemNum:    2,
				},
				giftItem: &GiftItem{
					SeedBase: SeedBase{ID: 10},
				},
				streamerResult: []RewardContent{
					{ContentType: CONTENT_TYPE_ITEM, ContentID: 1001, ContentQuantity: 1},
				},
				senderResult: []RewardContent{
					{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 1000},
				},
			},
			want: &GiftStockLog{
				MirrativID:           "user1",
				GiftTransactionLogID: 10,
				GiftTransactionLog: GiftTransactionLog{
					LogBase:    LogBase{gorm.Model{ID: 10}},
					StreamerID: "user1",
					ItemNum:    2,
				},
				GiftItemID: 10,
				GiftItem: GiftItem{
					SeedBase: SeedBase{ID: 10},
				},
				Stock: 1,
				Result: datatypes.NewJSONType(
					GiftStockResultLog{
						StreamerResults: []RewardContent{{ContentType: CONTENT_TYPE_ITEM, ContentID: 1001, ContentQuantity: 1}},
						SenderResults:   []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 1000}},
					},
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGachaGiftStockLog(tt.args.txLog, tt.args.giftItem, tt.args.streamerResult, tt.args.senderResult); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGachaGiftStockLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftStockLog_UpdateStock(t *testing.T) {
	type fields struct {
		LogBase              LogBase
		MirrativID           string
		GiftTransactionLogID uint
		GiftTransactionLog   GiftTransactionLog
		GiftItemID           uint
		GiftItem             GiftItem
		Stock                uint
		IsShowPopup          bool
		Result               datatypes.JSONType[GiftStockResultLog]
	}
	type args struct {
		stock uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GiftStockLog
	}{
		{
			name: "ストック更新",
			fields: fields{
				Stock:       3,
				IsShowPopup: true,
			},
			args: args{
				stock: 0,
			},
			want: &GiftStockLog{
				Stock:       0,
				IsShowPopup: true,
			},
		},
		{
			name: "ClientがGift通知をした",
			fields: fields{
				Stock:       3,
				IsShowPopup: false,
			},
			args: args{
				stock: 3,
			},
			want: &GiftStockLog{
				Stock:       3,
				IsShowPopup: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GiftStockLog{
				LogBase:              tt.fields.LogBase,
				MirrativID:           tt.fields.MirrativID,
				GiftTransactionLogID: tt.fields.GiftTransactionLogID,
				GiftTransactionLog:   tt.fields.GiftTransactionLog,
				GiftItemID:           tt.fields.GiftItemID,
				GiftItem:             tt.fields.GiftItem,
				Stock:                tt.fields.Stock,
				Result:               tt.fields.Result,
			}
			l.UpdateStock(tt.args.stock)
			if !reflect.DeepEqual(l, tt.want) {
				t.Errorf("GiftStockLog.UpdateStock() = %v, want %v", l, tt.want)
			}
		})
	}
}
