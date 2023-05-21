package main

// TODO: Configure linter imports
import (
	"fmt"

	"github.com/spf13/pflag"
)

const version = "0.1.0"
const usage = `mclogs is a fake log generator

Usage: mclogs [options]

Version: %s

Options:
  -f, --format string      log format. available formats:
                           - json (default)
                           - logfmt
  -o, --output string      output filename. Path-like is allowed. (default "stdout")
  -n, --number integer     number of lines to generate. (default "100")
  -b, --bytes integer      size of logs to generate (in bytes).
                           "bytes" will be ignored when "number" is set.
  -s, --sleep duration     fix creation time interval for each log (default unit "seconds"). It does not actually sleep.
                           examples: 10, 20ms, 5s, 1m
  -d, --delay duration     delay log generation speed (default unit "seconds").
                           examples: 10, 20ms, 5s, 1m
  -w, --overwrite          overwrite the existing log files.
  -l, --loop               loop output forever until killed.
`

func init() {
	pflag.Usage = printUsage
}

func printUsage() {
	fmt.Printf(usage, version)
}

func printVersion() {
	fmt.Printf("mclogs version %s\n", version)
}
