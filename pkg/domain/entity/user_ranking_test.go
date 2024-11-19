package entity

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUserRanking(t *testing.T) {
	type args struct {
		userID         uint
		rankingGroupID uint
	}
	tests := []struct {
		name string
		args args
		want *UserRanking
	}{
		{
			name: "UserRankingを生成できる",
			args: args{
				userID:         1,
				rankingGroupID: 1,
			},
			want: &UserRanking{
				UserID:         1,
				RankingGroupID: 1,
				Score:          0,
				ReceivedAt:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRanking(tt.args.userID, tt.args.rankingGroupID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRanking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRanking_AddScore(t *testing.T) {
	type fields struct {
		Score uint
	}
	type args struct {
		score uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "スコアが加算される",
			fields: fields{
				Score: 0,
			},
			args: args{
				score: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRanking{
				Score: tt.fields.Score,
			}
			u.AddScore(tt.args.score)

			if u.Score != tt.args.score+tt.fields.Score {
				t.Errorf("AddScore() = %v, want %v", u.Score, tt.args.score+tt.fields.Score)
			}
		})
	}
}

func TestUserRanking_UpdateBestScore(t *testing.T) {
	highScore := uint(1000)
	type fields struct {
		Score uint
	}
	type args struct {
		score uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "スコアが更新される",
			fields: fields{
				Score: 0,
			},
			args: args{
				score: highScore,
			},
			want: true,
		},
		{
			name: "スコアが更新されない",
			fields: fields{
				Score: highScore,
			},
			args: args{
				score: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRanking{
				Score: tt.fields.Score,
			}
			if got := u.UpdateBestScore(tt.args.score); got != tt.want {
				t.Errorf("UserRanking.UpdateBestScore() = %v, want %v", got, tt.want)
			}
			if u.Score != highScore {
				t.Errorf("UpdateBestScore() = %v, want %v", u.Score, highScore)
			}
		})
	}
}

func TestUserRanking_IsReceived(t *testing.T) {
	type fields struct {
		ReceivedAt *DateTime
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "受け取り済みの場合",
			fields: fields{
				ReceivedAt: &DateTime{},
			},
			want: true,
		},
		{
			name: "受け取り済みでない場合",
			fields: fields{
				ReceivedAt: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRanking{
				ReceivedAt: tt.fields.ReceivedAt,
			}
			if got := u.IsReceived(); got != tt.want {
				t.Errorf("UserRanking.IsReceived() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRanking_Receive(t *testing.T) {
	type fields struct {
		ReceivedAt *DateTime
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
			name: "受け取り日時が設定される",
			fields: fields{
				ReceivedAt: nil,
			},
			args: args{
				now: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserRanking{
				ReceivedAt: tt.fields.ReceivedAt,
			}
			u.Receive(tt.args.now)

			if u.ReceivedAt == nil {
				t.Errorf("Receive() = %v, want %v", u.ReceivedAt, tt.args.now)
			}
		})
	}
}
