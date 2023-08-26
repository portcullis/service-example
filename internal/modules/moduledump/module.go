package moduledump

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/portcullis/application"
)

type module int

// New returns information about the loaded modules
func New() module {
	return module(0)
}

func (m module) Start(ctx context.Context) error {
	app := application.FromContext(ctx)
	if app == nil {
		return nil
	}

	app.Controller.Range(func(name string, m application.Module) bool {
		slog.Debug(fmt.Sprintf("Module %q is %T", name, m))
		return true
	})

	return nil
}

func (m module) Stop(ctx context.Context) error {
	return nil
}
