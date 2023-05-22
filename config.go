package mclogs

import "time"

type Config struct {
	Format    Format
	Output    string
	Lines     int
	Bytes     int
	Sleep     time.Duration
	Delay     time.Duration
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
		Overwrite: false,
		Forever:   false,
	}
}
