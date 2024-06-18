//go:build !goverter

package app

import (
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"

	"go.uber.org/fx"
)

var Module = fx.Module("app",
	flow.Module,
	account.Module,
)
