//go:build !windows
// +build !windows

package mclogs

import (
	"errors"
	"os"
	"path/filepath"
	"syscall"
)

// Run checks overwrite flag and generates logs with given configuration
func Run(cfg Config) error {
	// If output is specified, prepare the file system
	// Otherwise, use stdout as output for the logs.
	if len(cfg.Output) > 0 {
		logDir := filepath.Dir(cfg.Output)

		if err := createLogsDir(logDir); err != nil {
			return err
		}

		if _, err := os.Stat(cfg.Output); err == nil && !cfg.Overwrite {
			return errors.New(cfg.Output + " already exists. You can overwrite with -w option")
		}
	}

	return Generate(cfg)
}

func createLogsDir(logsDir string) error {
	oldMask := syscall.Umask(0000)
	if err := os.MkdirAll(logsDir, 0766); err != nil {
		return err
	}
	syscall.Umask(oldMask)
	return nil
}
