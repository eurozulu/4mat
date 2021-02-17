package parser

import (
	"context"
	"github.com/eurozulu/4mat/formats"
	"log"
	"os"
	"sync"
)

const batchSize = 4

// ParseFiles scans the given set of filepaths for files which can be parsed into a known format.
// the resulting channel is the parsed form of the file, as defined by the format which Unmarshaled it.
func ParseFiles(ctx context.Context, filepaths ...string) <-chan interface{} {
	return parseFilePaths(ctx, ScanFilePaths(ctx, filepaths...))
}

func parseFilePaths(ctx context.Context, filePaths <-chan string) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)

		// Batch filepaths into small batches of batchSize to read concurrently
		var fps []string
		for {
			select {
			case <-ctx.Done():
				return
			case fp, ok := <-filePaths:
				if !ok {
					parseFileBatch(ctx, fps, ch)
					return
				}
				fps = append(fps, fp)
				if len(fps) >= batchSize {
					parseFileBatch(ctx, fps, ch)
					fps = nil
				}
			}
		}
	}()
	return ch
}

func parseFileBatch(ctx context.Context, fps []string, out chan<- interface{}) {
	var wg sync.WaitGroup
	wg.Add(len(fps))

	// Each routine writes result to its given pointer in the slice.
	// so results keep the same order as the filepaths.
	// No locking employed as each routine has unique pointer.
	result := make([]*interface{}, len(fps))
	for i, fp := range fps {
		go parseFile(fp, result[i], &wg)
	}
	wg.Wait()

	// When all done, send out the results in order, skipping those that failed.
	var index int
	for index < len(result) {
		if result[index] == nil {
			continue
		}
		select {
		case <-ctx.Done():
			return
		case out <- *result[index]:
		}
	}
}

func parseFile(p string, r *interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	f, err := DetectFormat(p)
	if err != nil {
		log.Println(err)
		return
	}

	fs := formats.NewFormatUnmarshaler(f)
	if fs == nil {
		log.Println("unknown format of %s", p)
		return
	}

	fl, err := os.Open(p)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := fl.Close(); err != nil {
			log.Println(err)
		}
	}()
	v, err := fs.Unmarshal(fl)
	if err != nil {
		log.Println(err)
		return
	}
	r = &v
}
