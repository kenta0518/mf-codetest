package repository

import (
	"context"

	"github.com/kenta0518/mf-codetest/pkg/domain/entity"
)

type User interface {
	Create(ctx context.Context, name string) (*entity.User, error)
}

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
