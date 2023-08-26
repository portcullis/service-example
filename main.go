package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/portcullis/application"
	"github.com/portcullis/service-example/internal/modules/failer"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

var (
	debugFlag = flag.Bool("debug", false, "enable debug output")
)

func main() {
	flag.Parse()

	var lvlr slog.LevelVar
	w := os.Stdout
	logger := slog.New(tint.NewHandler(colorable.NewColorable(w), &tint.Options{
		Level:   &lvlr,
		NoColor: !isatty.IsTerminal(w.Fd()),
	}))

	if *debugFlag {
		lvlr.Set(slog.LevelDebug)
	}

	// set the global
	slog.SetDefault(logger)

	if err := application.Bootstrap("sample", "0.0.0",
		application.WithConfigFile("./application.hcl"),
		application.WithLogger(logger),
		application.WithModule("Dumper", moduledump.New()),
		application.WithModule("Failer", failer.New()),
		application.WithModule("Simple", simple.New()),
	).Run(context.Background()); err != nil {
		panic(err)
	}
}
