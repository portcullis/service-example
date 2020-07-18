package moduledump

import (
	"context"

	"github.com/portcullis/logging"
	"github.com/portcullis/module"
)

type moduledumper int

// New returns information about the loaded modules
func New() module.Module {
	return new(moduledumper)
}

func (m *moduledumper) Start(ctx context.Context) error {
	controller := module.FromContext(ctx)
	if controller == nil {
		return nil
	}

	controller.Range(func(name string, m module.Module) bool {
		logging.Debug("Module %q is %T", name, m)
		return true
	})

	return nil
}

func (m *moduledumper) Stop(ctx context.Context) error {
	return nil
}
