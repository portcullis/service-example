package main

import (
	"github.com/portcullis/application"
	"github.com/portcullis/service-example/internal/modules/failer"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

func main() {
	if err := application.Run("sample", "v0.0.0",
		application.Module("Dumper", moduledump.New()),
		application.Module("Failer", failer.New()),
		application.Module("Simple", simple.New()),
	); err != nil {
		panic(err)
	}
}
