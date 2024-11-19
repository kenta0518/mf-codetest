package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserSoloRaid(t *testing.T) {
	type args struct {
		userID     uint
		soloRaidID uint
		check      time.Time
	}
	tests := []struct {
		name string
		args args
		want *UserSoloRaid
	}{
		{
			name: "UserSoloRaidの作成",
			args: args{
				userID:     1,
				soloRaidID: 1,
				check:      time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: &UserSoloRaid{
				UserResourceBase: UserResourceBase{
					UserID:     1,
					ResourceID: 1,
				},
				BossLevel:    1,
				Damage:       0,
				KillCount:    0,
				AllClearedAt: nil,
				Stamina:      10,
				CheckedAt:    DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserSoloRaid(tt.args.userID, tt.args.soloRaidID, tt.args.check); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserSoloRaid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSoloRaid_IsAllCleared(t *testing.T) {
	type fields struct {
		AllClearedAt *DateTime
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "AllClearedAtがnilの場合",
			fields: fields{
				AllClearedAt: nil,
			},
			want: false,
		},
		{
			name: "AllClearedAtがnilでない場合",
			fields: fields{
				AllClearedAt: &DateTime{Time: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserSoloRaid{
				AllClearedAt: tt.fields.AllClearedAt,
			}
			if got := u.IsAllCleared(); got != tt.want {
				t.Errorf("UserSoloRaid.IsAllCleared() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSoloRaid_AllCreal(t *testing.T) {
	type fields struct {
		AllClearedAt *DateTime
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "AllClearedAtに現在時刻を設定",
			fields: fields{
				AllClearedAt: nil,
			},
			args: args{
				now: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				AllClearedAt: tt.fields.AllClearedAt,
			}
			u.AllCreal(tt.args.now)

			if u.AllClearedAt == nil {
				t.Errorf("UserSoloRaid.AllCreal() = %v, want not nil", u.AllClearedAt)
			}
		})
	}
}

func TestUserSoloRaid_AddDamage(t *testing.T) {
	type fields struct {
		Damage uint
	}
	type args struct {
		damage uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Damageにダメージを追加",
			fields: fields{
				Damage: 0,
			},
			args: args{
				damage: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				Damage: tt.fields.Damage,
			}
			u.AddDamage(tt.args.damage)

			if u.Damage != tt.args.damage+tt.fields.Damage {
				t.Errorf("UserSoloRaid.AddDamage() = %v, want %v", u.Damage, tt.args.damage+u.Damage)
			}
		})
	}
}

func TestUserSoloRaid_LevelUp(t *testing.T) {
	type fields struct {
		BossLevel uint
		Damage    uint
		KillCount uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "BossLevel, KillCountを1増加",
			fields: fields{
				BossLevel: 1,
				Damage:    0,
				KillCount: 0,
			},
		},
		{
			name: "Damageを0にリセット",
			fields: fields{
				BossLevel: 1,
				Damage:    100,
				KillCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				BossLevel: tt.fields.BossLevel,
				Damage:    tt.fields.Damage,
				KillCount: tt.fields.KillCount,
			}
			u.LevelUp()

			if u.BossLevel != tt.fields.BossLevel+1 {
				t.Errorf("UserSoloRaid.LevelUp() = %v, want %v", u.BossLevel, tt.fields.BossLevel+1)
			}
			if u.KillCount != tt.fields.KillCount+1 {
				t.Errorf("UserSoloRaid.LevelUp() = %v, want %v", u.KillCount, tt.fields.KillCount+1)
			}
			if u.Damage != 0 {
				t.Errorf("UserSoloRaid.LevelUp() = %v, want %v", u.Damage, 0)
			}
		})
	}
}

func TestUserSoloRaid_ReceveryStamina(t *testing.T) {
	type fields struct {
		Stamina   uint
		CheckedAt DateTime
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want2  uint
		want3  time.Time
	}{
		{
			name: "Staminaが最大値の場合",
			fields: fields{
				Stamina:   SoloRaidStaminaMax,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want:  false,
			want2: SoloRaidStaminaMax,
			want3: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから10分経過していない",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 5, 0, 0, time.Local),
			},
			want:  false,
			want2: 5,
			want3: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから10分経過している",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 12, 0, 0, time.Local),
			},
			want:  true,
			want2: 6,
			want3: time.Date(2023, 5, 1, 0, 10, 0, 0, time.Local),
		},
		{
			name: "前回のチェックから20分経過している",
			fields: fields{
				Stamina:   5,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 28, 0, 0, time.Local),
			},
			want:  true,
			want2: 7,
			want3: time.Date(2023, 5, 1, 0, 20, 0, 0, time.Local),
		},
		{
			name: "Stamina最大値以上に回復しようとした場合",
			fields: fields{
				Stamina:   0,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				now: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
			},
			want:  true,
			want2: SoloRaidStaminaMax,
			want3: time.Date(2023, 5, 1, 12, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				Stamina:   tt.fields.Stamina,
				CheckedAt: tt.fields.CheckedAt,
			}
			if got := u.RecoveryStamina(tt.args.now); got != tt.want {
				t.Errorf("UserSoloRaid.ReceveryStamina() = %v, want %v", got, tt.want)
			}
			if u.Stamina != tt.want2 {
				t.Errorf("UserSoloRaid.ReceveryStamina() = %v, want %v", u.Stamina, tt.want2)
			}
			if u.CheckedAt.Time != tt.want3 {
				t.Errorf("UserSoloRaid.ReceveryStamina() = %v, want %v", u.CheckedAt.Time, tt.want3)
			}
		})
	}
}

func TestUserSoloRaid_CheckStamina(t *testing.T) {
	type fields struct {
		Stamina uint
	}
	type args struct {
		overAttack bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "overAttackがfalseの場合",
			fields: fields{
				Stamina: 10,
			},
			args: args{
				overAttack: false,
			},
			want: true,
		},
		{
			name: "overAttackがtrueの場合",
			fields: fields{
				Stamina: 10,
			},
			args: args{
				overAttack: true,
			},
			want: true,
		},
		{
			name: "スタミナが足りない場合",
			fields: fields{
				Stamina: 0,
			},
			args: args{
				overAttack: false,
			},
			want: false,
		},
		{
			name: "スタミナが足りない場合",
			fields: fields{
				Stamina: 2,
			},
			args: args{
				overAttack: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				Stamina: tt.fields.Stamina,
			}
			if got := u.CheckStamina(tt.args.overAttack); got != tt.want {
				t.Errorf("UserSoloRaid.CheckStamina() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserSoloRaid_UseStamina(t *testing.T) {
	type fields struct {
		Stamina   uint
		CheckedAt DateTime
	}
	type args struct {
		overAttack bool
		now        time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "overAttackがfalseの場合",
			fields: fields{
				Stamina:   9,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				overAttack: false,
				now:        time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "overAttackがtrueの場合",
			fields: fields{
				Stamina:   9,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				overAttack: true,
				now:        time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "スタミナが最大値の場合",
			fields: fields{
				Stamina:   10,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				overAttack: false,
				now:        time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "スタミナが足りない場合",
			fields: fields{
				Stamina:   0,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				overAttack: false,
				now:        time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
		{
			name: "スタミナが足りない場合",
			fields: fields{
				Stamina:   2,
				CheckedAt: DateTime{Time: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local)},
			},
			args: args{
				overAttack: true,
				now:        time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserSoloRaid{
				Stamina:   tt.fields.Stamina,
				CheckedAt: tt.fields.CheckedAt,
			}
			if got := u.UseStamina(tt.args.overAttack, tt.args.now); got != tt.want {
				t.Errorf("UserSoloRaid.UseStamina() = %v, want %v", got, tt.want)
			}

			if tt.fields.Stamina == SoloRaidStaminaMax && u.CheckedAt.Time != tt.args.now {
				t.Errorf("UserSoloRaid.UseStamina() = %v, want %v", u.CheckedAt.Time, tt.args.now)
			}
		})
	}
}
