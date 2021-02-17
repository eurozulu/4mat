package formats

import (
	"context"
	"io"
	"log"
)

// FormatWriter is a convienience struct to combine a Marshaler with an output stream.
// Calling its Write method, writes to the marshaller with the output stream.
// Also supports a chan feed of interfaces to write.
type FormatWriter struct {
	Marshaler FormatMarshaler
	Out       io.Writer
}

func (fw FormatWriter) Write(i interface{}) error {
	return fw.Marshaler.Marshal(i, fw.Out)
}

func (fw FormatWriter) WriteStream(ctx context.Context, src <-chan interface{}) <-chan bool {
	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case <-ctx.Done():
				return
			case i, ok := <-src:
				if !ok {
					return
				}
				if err := fw.Write(i); err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}()
	return done
}

type FormatUnmarshaler interface {
	Unmarshal(in io.Reader) (interface{}, error)
}

type FormatMarshaler interface {
	Marshal(v interface{}, out io.Writer) error
}

func NewFormatUnmarshaler(f Format) FormatUnmarshaler {
	switch f {
	case FormatYAML:
		return &FormatYaml{}
	case FormatJSON:
		return &FormatJson{}

	default:
		return nil
	}
}

func NewFormatMarshaler(f Format) FormatMarshaler {
	switch f {
	case FormatYAML:
		return &FormatYaml{}
	case FormatJSON:
		return &FormatJson{}

	default:
		return nil
	}
}
