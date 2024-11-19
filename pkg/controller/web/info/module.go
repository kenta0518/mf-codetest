package info

import "go.uber.org/fx"

func Modules() fx.Option {
	return fx.Module("info",
		fx.Provide(
			NewInfoController,
		),
	)
}
