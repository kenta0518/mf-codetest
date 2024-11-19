package entity

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewUserDeck(t *testing.T) {
	type args struct {
		userID           uint
		characterCardID1 uint
	}
	tests := []struct {
		name string
		args args
		want *UserDeck
	}{
		{
			name: "Instance作成",
			args: args{
				userID:           1000,
				characterCardID1: 1001,
			},
			want: &UserDeck{
				UserID:                   1000,
				MainUserCharacterCardID1: 1001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserDeck(tt.args.userID, tt.args.characterCardID1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDeck_SafeChange(t *testing.T) {
	type fields struct {
		Model                    gorm.Model
		UserID                   uint
		User                     User
		MainUserCharacterCardID1 uint
		MainUserCharacterCardID2 uint
		MainUserCharacterCardID3 uint
		MainUserCharacterCardID4 uint
		SubUserCharacterCardID1  uint
		SubUserCharacterCardID2  uint
		SubUserCharacterCardID3  uint
		SubUserCharacterCardID4  uint
	}
	type args struct {
		slotIndex   SlotIndex
		slotType    SlotType
		charaCardID uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Deckの変更",
			fields: fields{
				MainUserCharacterCardID1: 1000,
			},
			args: args{
				slotIndex:   SlotIndex1,
				slotType:    SlotTypeMain,
				charaCardID: 1001,
			},
			want: true,
		},
		{
			name: "カードの入れ替え",
			fields: fields{
				MainUserCharacterCardID1: 1000,
				MainUserCharacterCardID2: 1001,
			},
			args: args{
				slotIndex:   SlotIndex1,
				slotType:    SlotTypeMain,
				charaCardID: 1001,
			},
			want: true,
		},
		{
			name:   "SlotIndexの範囲外",
			fields: fields{},
			args: args{
				slotIndex: SlotIndex(5),
			},
			want: false,
		},
		{
			name:   "Main1のカードが必須",
			fields: fields{},
			args: args{
				slotIndex:   SlotIndex1,
				slotType:    SlotTypeMain,
				charaCardID: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UserDeck{
				Model:                    tt.fields.Model,
				UserID:                   tt.fields.UserID,
				User:                     tt.fields.User,
				MainUserCharacterCardID1: tt.fields.MainUserCharacterCardID1,
				MainUserCharacterCardID2: tt.fields.MainUserCharacterCardID2,
				MainUserCharacterCardID3: tt.fields.MainUserCharacterCardID3,
				MainUserCharacterCardID4: tt.fields.MainUserCharacterCardID4,
				SubUserCharacterCardID1:  tt.fields.SubUserCharacterCardID1,
				SubUserCharacterCardID2:  tt.fields.SubUserCharacterCardID2,
				SubUserCharacterCardID3:  tt.fields.SubUserCharacterCardID3,
				SubUserCharacterCardID4:  tt.fields.SubUserCharacterCardID4,
			}
			if got := d.SafeChange(tt.args.slotIndex, tt.args.slotType, tt.args.charaCardID); got != tt.want {
				t.Errorf("UserDeck.SafeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDeck_checkOrganizedCharaCard(t *testing.T) {
	type fields struct {
		Model                    gorm.Model
		UserID                   uint
		User                     User
		MainUserCharacterCardID1 uint
		MainUserCharacterCardID2 uint
		MainUserCharacterCardID3 uint
		MainUserCharacterCardID4 uint
		SubUserCharacterCardID1  uint
		SubUserCharacterCardID2  uint
		SubUserCharacterCardID3  uint
		SubUserCharacterCardID4  uint
	}
	type args struct {
		targetID uint
		changeID uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Main1のカードを入れ替え",
			fields: fields{
				MainUserCharacterCardID1: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Main2のカードを入れ替え",
			fields: fields{
				MainUserCharacterCardID2: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Main3のカードを入れ替え",
			fields: fields{
				MainUserCharacterCardID3: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Main4のカードを入れ替え",
			fields: fields{
				MainUserCharacterCardID4: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Sub1のカードを入れ替え",
			fields: fields{
				SubUserCharacterCardID1: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Sub2のカードを入れ替え",
			fields: fields{
				SubUserCharacterCardID2: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Sub3のカードを入れ替え",
			fields: fields{
				SubUserCharacterCardID3: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
		{
			name: "Sub4のカードを入れ替え",
			fields: fields{
				SubUserCharacterCardID4: 1000,
			},
			args: args{
				targetID: 1000,
				changeID: 1001,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UserDeck{
				Model:                    tt.fields.Model,
				UserID:                   tt.fields.UserID,
				User:                     tt.fields.User,
				MainUserCharacterCardID1: tt.fields.MainUserCharacterCardID1,
				MainUserCharacterCardID2: tt.fields.MainUserCharacterCardID2,
				MainUserCharacterCardID3: tt.fields.MainUserCharacterCardID3,
				MainUserCharacterCardID4: tt.fields.MainUserCharacterCardID4,
				SubUserCharacterCardID1:  tt.fields.SubUserCharacterCardID1,
				SubUserCharacterCardID2:  tt.fields.SubUserCharacterCardID2,
				SubUserCharacterCardID3:  tt.fields.SubUserCharacterCardID3,
				SubUserCharacterCardID4:  tt.fields.SubUserCharacterCardID4,
			}
			d.checkOrganizedCharaCard(tt.args.targetID, tt.args.changeID)
		})
	}
}

func TestUserDeck_getMainCharaCardID(t *testing.T) {
	type fields struct {
		Model                    gorm.Model
		UserID                   uint
		User                     User
		MainUserCharacterCardID1 uint
		MainUserCharacterCardID2 uint
		MainUserCharacterCardID3 uint
		MainUserCharacterCardID4 uint
		SubUserCharacterCardID1  uint
		SubUserCharacterCardID2  uint
		SubUserCharacterCardID3  uint
		SubUserCharacterCardID4  uint
	}
	type args struct {
		slotIndex SlotIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		{
			name: "Main1のカードを取得",
			fields: fields{
				MainUserCharacterCardID1: 1000,
			},
			args: args{
				slotIndex: SlotIndex1,
			},
			want: 1000,
		},
		{
			name: "Main2のカードを取得",
			fields: fields{
				MainUserCharacterCardID2: 1000,
			},
			args: args{
				slotIndex: SlotIndex2,
			},
			want: 1000,
		},
		{
			name: "Main3のカードを取得",
			fields: fields{
				MainUserCharacterCardID3: 1000,
			},
			args: args{
				slotIndex: SlotIndex3,
			},
			want: 1000,
		},
		{
			name: "Main4のカードを取得",
			fields: fields{
				MainUserCharacterCardID4: 1000,
			},
			args: args{
				slotIndex: SlotIndex4,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UserDeck{
				Model:                    tt.fields.Model,
				UserID:                   tt.fields.UserID,
				User:                     tt.fields.User,
				MainUserCharacterCardID1: tt.fields.MainUserCharacterCardID1,
				MainUserCharacterCardID2: tt.fields.MainUserCharacterCardID2,
				MainUserCharacterCardID3: tt.fields.MainUserCharacterCardID3,
				MainUserCharacterCardID4: tt.fields.MainUserCharacterCardID4,
				SubUserCharacterCardID1:  tt.fields.SubUserCharacterCardID1,
				SubUserCharacterCardID2:  tt.fields.SubUserCharacterCardID2,
				SubUserCharacterCardID3:  tt.fields.SubUserCharacterCardID3,
				SubUserCharacterCardID4:  tt.fields.SubUserCharacterCardID4,
			}
			if got := d.getMainCharaCardID(tt.args.slotIndex); got != tt.want {
				t.Errorf("UserDeck.getMainCharaCardID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDeck_getSubCharaCardID(t *testing.T) {
	type fields struct {
		Model                    gorm.Model
		UserID                   uint
		User                     User
		MainUserCharacterCardID1 uint
		MainUserCharacterCardID2 uint
		MainUserCharacterCardID3 uint
		MainUserCharacterCardID4 uint
		SubUserCharacterCardID1  uint
		SubUserCharacterCardID2  uint
		SubUserCharacterCardID3  uint
		SubUserCharacterCardID4  uint
	}
	type args struct {
		slotIndex SlotIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint
	}{
		{
			name: "Sub1のカードを取得",
			fields: fields{
				SubUserCharacterCardID1: 1000,
			},
			args: args{
				slotIndex: SlotIndex1,
			},
			want: 1000,
		},
		{
			name: "Sub2のカードを取得",
			fields: fields{
				SubUserCharacterCardID2: 1000,
			},
			args: args{
				slotIndex: SlotIndex2,
			},
			want: 1000,
		},
		{
			name: "Sub3のカードを取得",
			fields: fields{
				SubUserCharacterCardID3: 1000,
			},
			args: args{
				slotIndex: SlotIndex3,
			},
			want: 1000,
		},
		{
			name: "Sub4のカードを取得",
			fields: fields{
				SubUserCharacterCardID4: 1000,
			},
			args: args{
				slotIndex: SlotIndex4,
			},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &UserDeck{
				Model:                    tt.fields.Model,
				UserID:                   tt.fields.UserID,
				User:                     tt.fields.User,
				MainUserCharacterCardID1: tt.fields.MainUserCharacterCardID1,
				MainUserCharacterCardID2: tt.fields.MainUserCharacterCardID2,
				MainUserCharacterCardID3: tt.fields.MainUserCharacterCardID3,
				MainUserCharacterCardID4: tt.fields.MainUserCharacterCardID4,
				SubUserCharacterCardID1:  tt.fields.SubUserCharacterCardID1,
				SubUserCharacterCardID2:  tt.fields.SubUserCharacterCardID2,
				SubUserCharacterCardID3:  tt.fields.SubUserCharacterCardID3,
				SubUserCharacterCardID4:  tt.fields.SubUserCharacterCardID4,
			}
			if got := d.getSubCharaCardID(tt.args.slotIndex); got != tt.want {
				t.Errorf("UserDeck.getSubCharaCardID() = %v, want %v", got, tt.want)
			}
		})
	}
}
