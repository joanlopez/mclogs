package bootstrap

import (
	"fmt"

	"github.com/spf13/pflag"
)

const (
	version = "0.1.0"
	usage   = `mclogs is a fake log generator

Usage: mclogs [options]

Version: %s

Options:
  -f, --format string      log format - available formats:
                           - json (default)
                           - logfmt
  -o, --output string      output filename (default "stdout") - path-like is allowed 
  -n, --number integer     number of lines to generate (default "100")
  -b, --bytes integer      size of logs to generate (in bytes)
                           "bytes" will be ignored when "number" is set
  -s, --sleep duration     fix creation time interval for each log (default unit "seconds") - it does not actually sleep
                           examples: 10, 20ms, 5s, 1m
  -d, --delay duration     delay log generation speed (default unit "seconds")
                           examples: 10, 20ms, 5s, 1m
  -a, --at date            date and time -RFC3339- the log generation starts at (default "now")
                           examples: 2023-01-01T15:00:00+02:00
  -w, --overwrite          overwrite the existing log files
  -l, --loop               loop output forever until killed
`
)

func init() {
	pflag.Usage = printUsage
}

func printUsage() {
	fmt.Printf(usage, version)
}

func printVersion() {
	fmt.Printf("mclogs version %s\n", version)
}
