package failer

import (
	"context"
	"errors"
	"os"

	"github.com/portcullis/module"
)

type modulefailure struct {
}

// New returns module that will fail when env FAIL_MODE=1
func New() module.Module {
	if os.Getenv("FAIL_MODE") == "1" {
		return new(modulefailure)
	}
	return nil
}

func (m *modulefailure) Start(ctx context.Context) error {
	return errors.New("failer failed to start")
}

func (m *modulefailure) Stop(ctx context.Context) error {
	return errors.New("failer failed to stop")
}
