package couplefly

import (
	"flag"
	"fmt"
	"log/slog"
)

type Options struct {
	LogFormat string
	LogLevel  string
}

func Execute() int {
	options := parseArgs()

	logger, err := getLogger(options.LogLevel, options.LogFormat)
	if err != nil {
		slog.Error("getLogger", "error", err)
		return 1
	}

	slog.SetDefault(logger)

	err = run(options)
	if err != nil {
		slog.Error("run failed", "error", err)
		return 1
	}
	return 0
}

func parseArgs() Options {
	options := Options{}

	flag.StringVar(&options.LogLevel, "log-level", "info", "Log level (debug, info, warn, error), default: info")
	flag.StringVar(&options.LogFormat, "log-format", "text", "Log format (text or json)")

	flag.Parse()

	return options
}

func run(options Options) error {
	positionalArgs := flag.Args()

	fmt.Println("Named arguments:")
	fmt.Printf("log level: %s\n", options.LogLevel)
	fmt.Printf("log format: %s\n", options.LogFormat)

	if len(positionalArgs) > 0 {
		fmt.Println("Positional arguments:")
		for i, arg := range positionalArgs {
			fmt.Printf("%d: %s\n", i+1, arg)
		}
	}

	return nil
}
