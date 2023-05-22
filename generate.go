package mclogs

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Generate generates the logs with given configuration.
func Generate(cfg Config) error {
	writer, err := newWriter(cfg.Output)
	if err != nil {
		return err
	}

	defer func() {
		if writer != os.Stdout {
			_ = writer.Close()
			fmt.Println(cfg.Output, "is created.")
		}
	}()

	// TODO: Handle CTRL+C
	generateLoop(cfg, writer)

	return nil
}

func generateLoop(cfg Config, writer io.WriteCloser) {
	var (
		// TODO: Make 'at' configurable
		at      = time.Now()
		written struct {
			bytes int
			lines int
		}

		interval time.Duration
		delay    time.Duration
	)

	if cfg.Delay > 0 {
		interval = cfg.Delay
		delay = interval
	}

	if cfg.Sleep > 0 {
		interval = cfg.Sleep
	}

	for {
		if cfg.Bytes > 0 && written.bytes >= cfg.Bytes {
			return
		}

		if cfg.Lines > 0 && written.lines >= cfg.Lines {
			return
		}

		time.Sleep(delay)
		log := NewLog(cfg.Format, at)
		_, _ = writer.Write([]byte(log + "\n"))

		written.lines++
		written.bytes += len(log)
		at = at.Add(interval)
	}
}

// NewLog creates a log for given format.
func NewLog(format Format, at time.Time) string {
	return NewEvent(at).Format(format)
}

func newWriter(output string) (io.WriteCloser, error) {
	if len(output) == 0 {
		return os.Stdout, nil
	}

	logFile, err := os.Create(output)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}
