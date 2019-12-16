package main

import (
	"github.com/portcullis/application"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

func main() {
	if err := application.Run("sample", "v0.0.0",
		application.Module("Simple", simple.New()),
		application.Module("Dumper", moduledump.New()),
	); err != nil {
		panic(err)
	}
}
