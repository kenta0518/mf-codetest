package entity

import (
	"reflect"
	"testing"
)

func TestNewAdmin(t *testing.T) {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name string
		args args
		want *Admin
	}{
		{
			name: "Adminの生成",
			args: args{
				email:    "e-mainl",
				password: "password",
			},
			want: &Admin{
				Email:    "e-mainl",
				Password: "password",
				RoleType: RoleTypeMaster,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdmin(tt.args.email, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
