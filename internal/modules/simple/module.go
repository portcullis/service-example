package simple

import (
	"context"

	"github.com/portcullis/logging"
)

type module struct {
	cfg struct {
		Hello string `config:"hello,optional"`
	}
}

// New creates a simple module
func New() *module {
	return new(module)
}

func (m *module) Config() (interface{}, error) {
	logging.Debug("Config")
	return &m.cfg, nil
}

func (m *module) ConfigSet(v interface{}) error {
	logging.Debug("ConfigSet: %+v", v)
	return nil
}

func (m *module) Initialize(ctx context.Context) (context.Context, error) {
	logging.Debug("Initialize")
	return ctx, nil
}

func (m *module) PreStart(ctx context.Context) error {
	logging.Debug("PreStart")
	return nil
}

func (m *module) Start(ctx context.Context) error {
	logging.Debug("Start")
	logging.Info("Hello, %q!", m.cfg.Hello)
	return nil
}

func (m *module) PostStart(ctx context.Context) error {
	logging.Debug("PostStart")
	return nil
}

func (m *module) Stop(ctx context.Context) error {
	logging.Debug("Stop")
	return nil
}
