package entity

import (
	"reflect"
	"testing"
)

func TestNewUserMainAdv(t *testing.T) {
	type args struct {
		userId uint
		advID  uint
		isLock bool
	}
	tests := []struct {
		name string
		args args
		want *UserMainAdv
	}{
		{
			name: "UserAdvの生成",
			args: args{
				userId: 1000,
				advID:  1010,
				isLock: true,
			},
			want: &UserMainAdv{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1010,
				},
				IsLock: true,
			},
		},
		{
			name: "Unlockで生成",
			args: args{
				userId: 1000,
				advID:  1010,
				isLock: false,
			},
			want: &UserMainAdv{
				UserResourceBase: UserResourceBase{
					UserID:     1000,
					ResourceID: 1010,
				},
				IsLock: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserMainAdv(tt.args.userId, tt.args.advID, tt.args.isLock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAdv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserMainAdv_Unlock(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		MainAdv          MainAdv
		IsLock           bool
		IsReaded         bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "ロック解除",
			fields: fields{
				IsLock: true,
			},
		},
		{
			name: "ロック解除済み",
			fields: fields{
				IsLock: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UserMainAdv{
				UserResourceBase: tt.fields.UserResourceBase,
				MainAdv:          tt.fields.MainAdv,
				IsLock:           tt.fields.IsLock,
				IsReaded:         tt.fields.IsReaded,
			}
			a.Unlock()

			if a.IsLock {
				t.Errorf("Unlock() = %v, want %v", a.IsLock, false)
			}
		})
	}
}

func TestUserMainAdv_Read(t *testing.T) {
	type fields struct {
		UserResourceBase UserResourceBase
		MainAdv          MainAdv
		IsLock           bool
		IsReaded         bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "既読",
			fields: fields{
				IsReaded: false,
			},
		},
		{
			name: "既読済み",
			fields: fields{
				IsReaded: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UserMainAdv{
				UserResourceBase: tt.fields.UserResourceBase,
				MainAdv:          tt.fields.MainAdv,
				IsLock:           tt.fields.IsLock,
				IsReaded:         tt.fields.IsReaded,
			}
			a.Read()

			if !a.IsReaded {
				t.Errorf("Read() = %v, want %v", a.IsReaded, true)
			}
		})
	}
}
