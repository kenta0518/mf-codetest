package logger

import (
	"context"
	"time"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
)

type InventoryLog struct {
	ContentType     entity.ContentType `json:"content_type"`
	ContentID       uint               `json:"content_id"`
	ContentQuantity int                `json:"content_quantity"`
}

type CardSlot struct {
	CardID uint `json:"card_id"`
	Level  uint `json:"level"`
}

type EquipmentSlot struct {
	EquipmentID uint `json:"equipment_id"`
	Level       int  `json:"level"`
}

type Kpi interface {
	// インストールログ
	Install(ctx context.Context, userID uint, mirrativID string, userType string)
	// ログインログ
	Login(ctx context.Context, userID uint, mirrativID string, userType string, rank uint, rankExp uint, gold uint, inventory []InventoryLog, eventStamina uint, giftEventStamina uint)
	// アクセスログ
	Access(ctx context.Context, userID uint, mirrativID string, userType string, path string, installedAt time.Time, seconds uint, sessionID string)
	// バトル開始ログ
	StartBattle(ctx context.Context, userID uint, mirrativID string, stageID uint)
	// バトル終了ログ
	EndBattle(ctx context.Context, userID uint, mirrativID string, stageID uint, clear bool, score uint, clearMissionNum int)
	// Exバトル開始ログ
	StartExBattle(ctx context.Context, userID uint, mirrativID string, exStageID uint)
	// Exバトル終了ログ
	EndExBattle(ctx context.Context, userID uint, mirrativID string, exStageID uint, clear bool, score uint, clearMissionNum int)
	// ExバトルSkipログ
	SkipExBattle(ctx context.Context, userID uint, mirrativID string, exStageID uint, score uint, clearMissionNum int)
	// Eventバトル開始ログ
	StartEventBattle(ctx context.Context, userID uint, mirrativID string, eventStageID uint)
	// Eventバトル終了ログ
	EndEventBattle(ctx context.Context, userID uint, mirrativID string, eventStageID uint, clear bool, score uint, clearMissionNum int)
	// ちょこらんバトル開始ログ
	StartChokorankBattle(ctx context.Context, userID uint, mirrativID string, chokorankStageID uint)
	// ちょこらんバトル終了ログ
	EndChokorankBattle(ctx context.Context, userID uint, mirrativID string, chokorankStageID uint, clear bool, score uint, clearMissionNum int)
	// ギフト送信ログ
	SendGift(ctx context.Context, senderMirrativID string, streamerMirrativID string, giftItemID uint, quantity uint, mirrativCoin uint)
	// ハッピーボックス送信ログ
	SendHappyBox(ctx context.Context, senderMirrativID string, streamerMirrativID string, happyBoxID uint, quantity int)
	// 消費ログ
	Consume(ctx context.Context, userID uint, mirrativID string, contentType entity.ContentType, contentID uint, contentQuantity int, transactionID string, source string)
	// 獲得ログ
	Accure(ctx context.Context, userID uint, mirrativID string, contentType entity.ContentType, contentID uint, contentQuantity int, transactionID string, source string)
	// キャラクター強化ログ
	EnhanceCharacter(ctx context.Context, userID uint, mirrativID string, characterID uint, beforeLv uint, afterLv uint)
	// カード強化ログ
	EnhanceCard(ctx context.Context, userID uint, mirrativID string, cardID uint, beforeLv uint, afterLv uint)
	// カード覚醒ログ
	EvolveCard(ctx context.Context, userID uint, mirrativID string, uniqID uint, beforeID uint, afterID uint)
	// 装備強化ログ
	EnhanceEquipment(ctx context.Context, userID uint, mirrativID string, equipmentID uint, beforeID uint, afterLv uint)
	// 装備進化ログ
	EvolveEquipment(ctx context.Context, userID uint, mirrativID string, equipmentID uint, beforeLv uint, afterLv uint)
	// Adv再生ログ
	ReadAdv(ctx context.Context, userID uint, mirrativID string, advID uint)
	// 過去Adv再生ログ
	ReadPastAdv(ctx context.Context, userID uint, mirrativID string, advID uint)
	// ミッションログ
	ClearMission(ctx context.Context, userID uint, mirrativID string, missionID uint)
	// ガチャログ
	Gacha(ctx context.Context, userID uint, mirrativID string, gachaID uint, times int)
	// 交換所ログ
	Exchange(ctx context.Context, userID uint, mirrativID string, exchangeID uint, exchangeLineupID uint, quantity int)
	// 書き込み
	Flush() error
}

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
