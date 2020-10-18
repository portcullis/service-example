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
		application.WithConfigFile("./application.hcl"),

		application.WithModule("Dumper", moduledump.New()),
		application.WithModule("Failer", failer.New()),
		application.WithModule("Simple", simple.New()),
	).Run(context.Background()); err != nil {
		panic(err)
	}
}
