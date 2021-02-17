package main

import (
	"context"
	"fmt"
	"github.com/eurozulu/4mat/formats"
	"github.com/eurozulu/4mat/parser"
	"os"
)

func ToYaml(filepaths ...string) error {
	return StreamPaths(formats.FormatYAML, filepaths...)
}

type JsonCommand struct {
	PrettyPrint bool `flag:"prettyprint,pretty,pp"`
}

func (jc JsonCommand) ToJson(filepaths ...string) error {
	return StreamPaths(formats.FormatJSON, filepaths...)
}

func StreamPaths(format formats.Format, filepaths ...string) error {
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()

	outF := formats.NewFormatMarshaler(format)
	if outF == nil {
		return fmt.Errorf("unknown format")
	}
	out := &formats.FormatWriter{
		Marshaler: outF,
		Out:       os.Stdout,
	}

	done := out.WriteStream(ctx, parser.ParseFiles(ctx, filepaths...))

	<-done
	return nil
}
