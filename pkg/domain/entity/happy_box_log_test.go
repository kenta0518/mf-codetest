package entity

import (
	"reflect"
	"testing"

	"gorm.io/datatypes"
)

func TestNewSenderHappyBoxLog(t *testing.T) {
	type args struct {
		userId        uint
		content       RewardContent
		result        []RewardContent
		transactionID string
	}
	tests := []struct {
		name string
		args args
		want *HappyBoxLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userId:        1,
				content:       RewardContent{ContentType: CONTENT_TYPE_ITEM, ContentID: 1001, ContentQuantity: 1},
				result:        []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
				transactionID: "transaction",
			},
			want: &HappyBoxLog{
				UserID:          1,
				IsSender:        true,
				ContentType:     CONTENT_TYPE_ITEM,
				ContentID:       1001,
				ContentQuantity: 1,
				TransactionID:   "transaction",
				Result: datatypes.NewJSONType(HappyBoxResultLog{
					Results: []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSenderHappyBoxLog(tt.args.userId, tt.args.content, tt.args.result, tt.args.transactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHappyBoxLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStreamerHappyBoxLog(t *testing.T) {
	type args struct {
		userId        uint
		content       RewardContent
		result        []RewardContent
		transactionID string
	}
	tests := []struct {
		name string
		args args
		want *HappyBoxLog
	}{
		{
			name: "インスタンス作れるか",
			args: args{
				userId:        1,
				content:       RewardContent{ContentType: CONTENT_TYPE_ITEM, ContentID: 1001, ContentQuantity: 1},
				result:        []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
				transactionID: "transaction",
			},
			want: &HappyBoxLog{
				UserID:          1,
				IsSender:        false,
				ContentType:     CONTENT_TYPE_ITEM,
				ContentID:       1001,
				ContentQuantity: 1,
				TransactionID:   "transaction",
				Result: datatypes.NewJSONType(HappyBoxResultLog{
					Results: []RewardContent{{ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 100}},
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStreamerHappyBoxLog(tt.args.userId, tt.args.content, tt.args.result, tt.args.transactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStreamerHappyBoxLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
