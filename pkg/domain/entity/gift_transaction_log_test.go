package entity

import (
	"reflect"
	"testing"
)

func TestGiftTransactionLog_Close(t *testing.T) {
	type fields struct {
		LogBase          LogBase
		ItemID           uint
		ItemNum          uint
		StreamerID       string
		StreamerName     string
		StreamerImageUrl string
		SenderID         string
		SenderName       string
		SenderImageUrl   string
		LiveID           string
		TransactionID    string
		OpenedAt         string
		State            GiftTransactionState
	}
	tests := []struct {
		name   string
		fields fields
		want   *GiftTransactionLog
	}{
		{
			name: "Closeできるか",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_OPEN,
			},
			want: &GiftTransactionLog{State: GIFT_TRANSACTION_STATE_CLOSE},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &GiftTransactionLog{
				LogBase:          tt.fields.LogBase,
				ItemID:           tt.fields.ItemID,
				ItemNum:          tt.fields.ItemNum,
				StreamerID:       tt.fields.StreamerID,
				StreamerName:     tt.fields.StreamerName,
				StreamerImageUrl: tt.fields.StreamerImageUrl,
				SenderID:         tt.fields.SenderID,
				SenderName:       tt.fields.SenderName,
				SenderImageUrl:   tt.fields.SenderImageUrl,
				LiveID:           tt.fields.LiveID,
				TransactionID:    tt.fields.TransactionID,
				OpenedAt:         tt.fields.OpenedAt,
				State:            tt.fields.State,
			}
			log.Close()
			if !reflect.DeepEqual(log, tt.want) {
				t.Errorf("GiftTransactionLog.Close() = %v, want %v", log, tt.want)
			}
		})
	}
}

func TestGiftTransactionLog_Cancel(t *testing.T) {
	type fields struct {
		LogBase          LogBase
		ItemID           uint
		ItemNum          uint
		StreamerID       string
		StreamerName     string
		StreamerImageUrl string
		SenderID         string
		SenderName       string
		SenderImageUrl   string
		LiveID           string
		TransactionID    string
		OpenedAt         string
		State            GiftTransactionState
	}
	tests := []struct {
		name   string
		fields fields
		want   *GiftTransactionLog
	}{
		{
			name: "Cancelできるか",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_OPEN,
			},
			want: &GiftTransactionLog{State: GIFT_TRANSACTION_STATE_CANCEL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &GiftTransactionLog{
				LogBase:          tt.fields.LogBase,
				ItemID:           tt.fields.ItemID,
				ItemNum:          tt.fields.ItemNum,
				StreamerID:       tt.fields.StreamerID,
				StreamerName:     tt.fields.StreamerName,
				StreamerImageUrl: tt.fields.StreamerImageUrl,
				SenderID:         tt.fields.SenderID,
				SenderName:       tt.fields.SenderName,
				SenderImageUrl:   tt.fields.SenderImageUrl,
				LiveID:           tt.fields.LiveID,
				TransactionID:    tt.fields.TransactionID,
				OpenedAt:         tt.fields.OpenedAt,
				State:            tt.fields.State,
			}
			log.Cancel()
			if !reflect.DeepEqual(log, tt.want) {
				t.Errorf("GiftTransactionLog.Cancel() = %v, want %v", log, tt.want)
			}
		})
	}
}

func TestGiftTransactionLog_IsCanceled(t *testing.T) {
	type fields struct {
		LogBase          LogBase
		ItemID           uint
		ItemNum          uint
		StreamerID       string
		StreamerName     string
		StreamerImageUrl string
		SenderID         string
		SenderName       string
		SenderImageUrl   string
		LiveID           string
		TransactionID    string
		OpenedAt         string
		State            GiftTransactionState
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Openならfalse",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_OPEN,
			},
			want: false,
		},
		{
			name: "Cancelならtrue",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_CANCEL,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &GiftTransactionLog{
				LogBase:          tt.fields.LogBase,
				ItemID:           tt.fields.ItemID,
				ItemNum:          tt.fields.ItemNum,
				StreamerID:       tt.fields.StreamerID,
				StreamerName:     tt.fields.StreamerName,
				StreamerImageUrl: tt.fields.StreamerImageUrl,
				SenderID:         tt.fields.SenderID,
				SenderName:       tt.fields.SenderName,
				SenderImageUrl:   tt.fields.SenderImageUrl,
				LiveID:           tt.fields.LiveID,
				TransactionID:    tt.fields.TransactionID,
				OpenedAt:         tt.fields.OpenedAt,
				State:            tt.fields.State,
			}
			if got := log.IsCanceled(); got != tt.want {
				t.Errorf("GiftTransactionLog.IsCanceled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGiftTransactionLog_IsClosed(t *testing.T) {
	type fields struct {
		LogBase          LogBase
		ItemID           uint
		ItemNum          uint
		StreamerID       string
		StreamerName     string
		StreamerImageUrl string
		SenderID         string
		SenderName       string
		SenderImageUrl   string
		LiveID           string
		TransactionID    string
		OpenedAt         string
		State            GiftTransactionState
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Openならfalse",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_OPEN,
			},
			want: false,
		},
		{
			name: "Closeならtrue",
			fields: fields{
				State: GIFT_TRANSACTION_STATE_CLOSE,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := &GiftTransactionLog{
				LogBase:          tt.fields.LogBase,
				ItemID:           tt.fields.ItemID,
				ItemNum:          tt.fields.ItemNum,
				StreamerID:       tt.fields.StreamerID,
				StreamerName:     tt.fields.StreamerName,
				StreamerImageUrl: tt.fields.StreamerImageUrl,
				SenderID:         tt.fields.SenderID,
				SenderName:       tt.fields.SenderName,
				SenderImageUrl:   tt.fields.SenderImageUrl,
				LiveID:           tt.fields.LiveID,
				TransactionID:    tt.fields.TransactionID,
				OpenedAt:         tt.fields.OpenedAt,
				State:            tt.fields.State,
			}
			if got := log.IsClosed(); got != tt.want {
				t.Errorf("GiftTransactionLog.IsClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}
