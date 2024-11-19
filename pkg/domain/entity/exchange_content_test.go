package entity

import (
	"reflect"
	"testing"
)

func TestExchangeContent_RewadContent(t *testing.T) {
	type fields struct {
		SeedBase         SeedBase
		ExchangeLineupID uint
		ContentType      ContentType
		ContentID        uint
		ContentQuantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *RewardContent
	}{
		{
			name: "RewardContentを返せるか",
			fields: fields{
				ContentType: CONTENT_TYPE_GOLD, ContentID: 0, ContentQuantity: 1000,
			},
			want: &RewardContent{
				CONTENT_TYPE_GOLD, 0, 1000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ExchangeContent{
				SeedBase:         tt.fields.SeedBase,
				ExchangeLineupID: tt.fields.ExchangeLineupID,
				ContentType:      tt.fields.ContentType,
				ContentID:        tt.fields.ContentID,
				ContentQuantity:  tt.fields.ContentQuantity,
			}
			if got := e.RewadContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeContent.RewadContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
