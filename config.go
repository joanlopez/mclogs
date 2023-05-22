package mclogs

import "time"

type Config struct {
	Format    Format
	Output    string
	Lines     int
	Bytes     int
	Sleep     time.Duration
	Delay     time.Duration
	At        time.Time
	Overwrite bool
	Forever   bool
}

func DefaultConfig() Config {
	return Config{
		Format:    FormatJSON,
		Output:    "",
		Lines:     100,
		Bytes:     0,
		Sleep:     0,
		Delay:     0,
		At:        time.Now(),
		Overwrite: false,
		Forever:   false,
	}
}
