package repository

import (
	"context"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
)

type Transaction interface {
	Create(ctx context.Context, userID int, amount int, description string) (*entity.Transaction, error)
	GetUserTotalAmountForUpdate(ctx context.Context, userID int) (int, error)
}

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
