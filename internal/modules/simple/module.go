package simple

import (
	"context"

	"github.com/portcullis/logging"
)

type module struct {
}

// New creates a simple module
func New() *module {
	return new(module)
}

func (s *module) Initialize(ctx context.Context) (context.Context, error) {
	logging.Debug("Initialize")
	return ctx, nil
}

func (s *module) PreStart(ctx context.Context) error {
	logging.Debug("PreStart")
	return nil
}

func (s *module) Start(ctx context.Context) error {
	logging.Debug("Start")
	return nil
}

func (s *module) PostStart(ctx context.Context) error {
	logging.Debug("PostStart")
	return nil
}

func (s *module) Stop(ctx context.Context) error {
	logging.Debug("Stop")
	return nil
}
