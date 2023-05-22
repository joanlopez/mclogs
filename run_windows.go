package mclogs

import (
	"errors"
	"os"
	"path/filepath"
)

// Run checks overwrite flag and generates logs with given configuration
func Run(cfg Config) error {
	// If output is specified, prepare the file system
	// Otherwise, use stdout as output for the logs.
	if len(cfg.Output) > 0 {
		logsDir := filepath.Dir(cfg.Output)

		if err := os.MkdirAll(logsDir, 0766); err != nil {
			return err
		}

		if _, err := os.Stat(cfg.Output); err == nil && !cfg.Overwrite {
			return errors.New(cfg.Output + " already exists. You can overwrite with -w option")
		}
	}

	return Generate(cfg)
}
