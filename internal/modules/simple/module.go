package simple

import (
	"context"

	"github.com/portcullis/logging"
	"github.com/portcullis/module"
)

type simple struct {
}

// New creates a simple module
func New() module.Module {
	return new(simple)
}

func (s *simple) Initialize(ctx context.Context) (context.Context, error) {
	logging.Debug("Initialize")
	return ctx, nil
}

func (s *simple) PreStart(ctx context.Context) error {
	logging.Debug("PreStart")
	return nil
}

func (s *simple) Start(ctx context.Context) error {
	logging.Debug("Start")
	return nil
}

func (s *simple) PostStart(ctx context.Context) error {
	logging.Debug("PostStart")
	return nil
}

func (s *simple) Stop(ctx context.Context) error {
	logging.Debug("Stop")
	return nil
}
