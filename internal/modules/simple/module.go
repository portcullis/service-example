package simple

import (
	"context"
	"log/slog"
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
	slog.Debug("Config")
	return &m.cfg, nil
}

func (m *module) ConfigSet(v interface{}) error {
	slog.Debug("ConfigSet", "val", v)
	return nil
}

func (m *module) Initialize(ctx context.Context) (context.Context, error) {
	slog.Debug("Initialize")
	return ctx, nil
}

func (m *module) PreStart(ctx context.Context) error {
	slog.Debug("PreStart")
	return nil
}

func (m *module) Start(ctx context.Context) error {
	slog.Debug("Start")
	slog.Info("Hello", "name", m.cfg.Hello)
	return nil
}

func (m *module) PostStart(ctx context.Context) error {
	slog.Debug("PostStart")
	return nil
}

func (m *module) Stop(ctx context.Context) error {
	slog.Debug("Stop")
	return nil
}
