package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/portcullis/logging"
	"github.com/portcullis/module"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

func main() {
	ctr := module.Controller{}
	ctr.Add("Simple", simple.New())
	ctr.Add("Dumper", moduledump.New())

	schan := make(chan os.Signal, 1)
	defer close(schan)

	signal.Notify(schan, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(schan)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := ctr.Run(ctx); err != nil && err != context.Canceled {
			logging.Error("Failed to run: %v", err)
		}
		wg.Done()
		cancel()
	}()

	select {
	case <-ctx.Done():

	case sig := <-schan:
		logging.Info("Signal %v received", sig)
		cancel()
	}

	wg.Wait()
}
