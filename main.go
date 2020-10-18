package main

import (
	"context"

	"github.com/portcullis/application"
	"github.com/portcullis/service-example/internal/modules/failer"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

func main() {
	if err := application.Bootstrap("sample", "0.0.0",
		application.RunModule("Dumper", moduledump.New()),
		application.RunModule("Failer", failer.New()),
		application.RunModule("Simple", simple.New()),
	).Run(context.Background()); err != nil {
		panic(err)
	}
}
