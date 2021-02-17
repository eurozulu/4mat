package formats

import (
	"fmt"
)

const (
	FormatUnknown = iota
	FormatYAML
	FormatJSON
	FormatXML
	FormatCSV
	FormatPDF
)

var Formats = []string{"unknown", "yaml", "json", "xml", "csv", "pdf"}

type Format int

func (f Format) String() string {
	t, err := f.MarshalText()
	if err != nil {
		return ""
	}
	return string(t)
}

func (f Format) MarshalText() (text []byte, err error) {
	if int(f) < 0 || int(f) >= len(Formats) {
		return nil, fmt.Errorf("format (%d) is not a known format", f)
	}
	return []byte(Formats[f]), nil
}

func (f *Format) UnmarshalText(text []byte) error {
	i := stringIndex(string(text), Formats)
	if i < 0 {
		return fmt.Errorf("Failed to parse %s as a known format", text)
	}
	*f = Format(i)
	return nil
}

func stringIndex(s string, ss []string) int {
	for i, sz := range ss {
		if s == sz {
			return i
		}
	}
	return -1
}
