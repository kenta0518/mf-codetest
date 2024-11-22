package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
	"github.com/kenta0518/mf-codetest/pkg/domain/repository"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) repository.Transaction {
	return &transactionRepository{
		db: db,
	}
}

func (t *transactionRepository) Create(ctx context.Context, userID int, amount int, description string) (*entity.Transaction, error) {
	tx, ok := GetTx(ctx)

	if !ok {
		return nil, repository.ErrTx
	}

	transaction := entity.Transaction{UserID: userID, Amount: amount, Description: description}
	err := tx.Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, err
}

func (t *transactionRepository) GetUserTotalAmountForUpdate(ctx context.Context, userID int) (int, error) {
	tx, ok := GetTx(ctx)

	// トランザクションが存在しない場合は、デフォルトのDB接続を使用
	if !ok {
		tx = t.db
	}

	var totalAmount sql.NullInt64
	// ユーザーごとの総取引金額をロックして取得
	err := tx.WithContext(ctx).Raw(`
		SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = ? FOR UPDATE
	`, userID).Scan(&totalAmount).Error
	if err != nil {
		return 0, fmt.Errorf("failed to get total amount for user %d: %w", userID, err)
	}

	// totalAmount.Valid が true の場合、値を返す
	if totalAmount.Valid {
		return int(totalAmount.Int64), nil
	}

	// totalAmount.Valid が false の場合は、0 を返す
	return 0, nil
}
