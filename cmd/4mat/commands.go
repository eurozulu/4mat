package main

import (
	"context"
	"fmt"
	"github.com/eurozulu/4mat/formats"
	"github.com/eurozulu/4mat/parser"
	"io"
	"os"
	"os/signal"
)

// Verbose displays file error logs usually surpressed. (Such as unknown format)
// TODO: Implement this with a common error handler for all the log.Println(err)
var Verbose bool

func ToYaml(filepaths ...string) error {
	return StreamPaths(os.Stdout, formats.FormatYAML, filepaths...)
}

type JsonCommand struct {
	PrettyPrint bool `flag:"prettyprint,pretty,pp"`
}

func (jc JsonCommand) ToJson(filepaths ...string) error {
	return StreamPaths(os.Stdout, formats.FormatJSON, filepaths...)
}

func StreamPaths(out io.Writer, format formats.Format, filepaths ...string) error {
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()

	outF := formats.NewFormatMarshaler(format)
	if outF == nil {
		return fmt.Errorf("unknown format")
	}
	fw := &formats.FormatWriter{
		Marshaler: outF,
		Out:       out,
	}

	done := fw.WriteStream(ctx, parser.ParseFiles(ctx, filepaths...))

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	case <-sig:
		return nil
	}
}
