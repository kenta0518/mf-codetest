package logger

import "context"

type KpiFactory interface {
	Create(ctx context.Context) (Kpi, error)
}

//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../../mock/$GOPACKAGE/$GOFILE
