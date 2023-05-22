package mclogs

import (
	"errors"
	"strings"
)

var (
	FormatJSON   = Format("json")
	FormatLogfmt = Format("logfmt")
)

var ErrFormatNotSupported = errors.New("format not supported")

type Format string

func NewFormat(f string) (Format, error) {
	// Clean up
	f = strings.TrimSpace(f)
	f = strings.ToLower(f)

	// Validate
	switch f {
	case FormatJSON.String():
		return FormatJSON, nil
	case FormatLogfmt.String():
		return FormatLogfmt, nil
	default:
		return "", ErrFormatNotSupported
	}
}

func (f Format) String() string {
	return string(f)
}
