package main

import (
	"flag"
	"os"

	"github.com/portcullis/application"
	"github.com/portcullis/config"
	"github.com/portcullis/logging"
	writer "github.com/portcullis/logging/format/simple"
	"github.com/portcullis/service-example/internal/modules/moduledump"
	"github.com/portcullis/service-example/internal/modules/simple"
)

func init() {
	// global logging setup
	level := logging.LevelInformational
	simpleWriter := writer.New(
		os.Stdout,
		writer.Level(level),
	)
	logging.DefaultLog = logging.New(
		logging.WithWriter(simpleWriter),
	)

	// create a setting
	levelSetting := &config.Setting{
		Name:         "Level",
		Path:         "Logging",
		Description:  "Logging level to output",
		DefaultValue: level.String(),
		Value:        &level,
	}

	// update the current log level
	levelSetting.Notify(config.NotifyFunc(func(s *config.Setting) {
		writer.Level(level)(simpleWriter)
	}))

	// make it a flag -level
	levelSetting.Flag("level", flag.CommandLine)
}

func main() {
	flag.Parse()

	if err := application.Run("sample", "v0.0.0",
		application.Module("Dumper", moduledump.New()),
		//application.Module("Failer", failer.New()),
		application.Module("Simple", simple.New()),
	); err != nil {
		panic(err)
	}
}
