package usecase

import (
	"context"
	"fmt"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
	"github.com/kenta0518/mf-codetest/pkg/domain/repository"
	"github.com/kenta0518/mf-codetest/pkg/usecase/model"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Transaction interface {
	CreateTransaction(ctx context.Context, userID int, amount int, description string) (*model.Transaction, error)
}

type transactionUsecase struct {
	TransactionRepository repository.Transaction
	localizer             *i18n.Localizer
	tx                    repository.DbTransaction
}

func NewTransactionUsecase(ua repository.Transaction, localizer *i18n.Localizer, tx repository.DbTransaction) Transaction {
	return &transactionUsecase{
		TransactionRepository: ua,
		localizer:             localizer,
		tx:                    tx,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, userID int, amount int, description string) (*model.Transaction, error) {
	value, err := u.tx.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		// 初期化
		totalAmount := 0

		// ユーザーの総取引金額をロックして取得
		totalAmount, err := u.TransactionRepository.GetUserTotalAmountForUpdate(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("failed to get total amount for user: %w", err)
		}

		// 合計金額が登録可能な上限を超えていないかをチェック
		const amountLimit = 1000
		if totalAmount+amount > amountLimit {
			c := &i18n.LocalizeConfig{MessageID: model.E0201}
			return nil, model.NewErrPaymentRequired(model.E0201, u.localizer.MustLocalize(c))
		}

		// 新しい取引を作成
		txn, err := u.TransactionRepository.Create(ctx, userID, amount, description)
		if err != nil {
			return nil, fmt.Errorf("failed to create transaction: %w", err)
		}

		return txn, nil
	})

	if err != nil {
		return nil, err
	}

	return model.NewTransaction(value.(*entity.Transaction)), nil
}
