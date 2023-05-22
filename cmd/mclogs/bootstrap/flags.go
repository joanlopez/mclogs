package bootstrap

import (
	"os"
	"time"

	"github.com/spf13/pflag"

	"github.com/joanlopez/mclogs"
)

const (
	timeFormat = time.RFC3339
)

func configFromFlags() mclogs.Config {
	cfg := mclogs.DefaultConfig()

	helpF := pflag.BoolP("help", "h", false, "")
	versionF := pflag.BoolP("version", "v", false, "")
	formatF := pflag.StringP("format", "f", cfg.Format.String(), "")
	outputF := pflag.StringP("output", "o", cfg.Output, "")
	linesF := pflag.IntP("number", "n", cfg.Lines, "")
	bytesF := pflag.IntP("bytes", "b", cfg.Bytes, "")
	sleepF := pflag.StringP("sleep", "s", "0s", "")
	delayF := pflag.StringP("delay", "d", "0s", "")
	atF := pflag.StringP("at", "a", cfg.At.Format(timeFormat), "")
	overwriteF := pflag.BoolP("overwrite", "w", false, "")
	foreverF := pflag.BoolP("loop", "l", false, "")

	pflag.Parse()

	if *helpF {
		printUsage()
		os.Exit(0)
	}

	if *versionF {
		printVersion()
		os.Exit(0)
	}

	captureFormat(&cfg, *formatF)
	captureLines(&cfg, *linesF)
	captureBytes(&cfg, *bytesF)
	captureSleep(&cfg, *sleepF)
	captureDelay(&cfg, *delayF)
	captureAt(&cfg, *atF)
	captureOutput(&cfg, *outputF)
	captureOverwrite(&cfg, *overwriteF)
	captureForever(&cfg, *foreverF)

	return cfg
}

func captureFormat(cfg *mclogs.Config, format string) {
	var err error
	if cfg.Format, err = mclogs.NewFormat(format); err != nil {
		panic(err)
	}
}

func captureLines(cfg *mclogs.Config, lines int) {
	if lines < 0 {
		panic("lines can not be negative")
	}

	cfg.Lines = lines
}

func captureBytes(cfg *mclogs.Config, bytes int) {
	if bytes < 0 {
		panic("bytes can not be negative")
	}

	cfg.Bytes = bytes
}

func captureSleep(cfg *mclogs.Config, sleep string) {
	var err error
	if cfg.Sleep, err = time.ParseDuration(sleep); err != nil {
		panic(err)
	}
}

func captureDelay(cfg *mclogs.Config, delay string) {
	var err error
	if cfg.Delay, err = time.ParseDuration(delay); err != nil {
		panic(err)
	}
}

func captureAt(cfg *mclogs.Config, at string) {
	var err error
	if cfg.At, err = time.Parse(timeFormat, at); err != nil {
		panic(err)
	}
}

func captureForever(cfg *mclogs.Config, forever bool) {
	if forever {
		cfg.Lines = 0
		cfg.Bytes = 0
	}
	cfg.Forever = forever
}

func captureOverwrite(cfg *mclogs.Config, overwrite bool) {
	cfg.Overwrite = overwrite
}

func captureOutput(cfg *mclogs.Config, output string) {
	cfg.Output = output
}
