package failer

import (
	"context"
	"errors"
	"os"
)

type module struct {
}

// New returns module that will fail when env FAIL_MODE=1
func New() *module {
	if os.Getenv("FAIL_MODE") == "1" {
		return new(module)
	}

	return nil
}

func (m *module) Start(ctx context.Context) error {
	return errors.New("failer failed to start")
}

func (m *module) Stop(ctx context.Context) error {
	return errors.New("failer failed to stop")
}
