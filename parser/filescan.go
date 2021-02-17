package parser

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

var FlagRecursive bool

// ScanFilePaths scans a given set of filepaths for files.
// All paths must point to a file.  A directory path throws an error unless recursive is true.
// If Recursive, true, file paths are treated the same, directories are listed, files first,
// followed by the files, if any, it ints subdirectories.
func ScanFilePaths(ctx context.Context, filepaths ...string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		for _, fp := range filepaths {
			rfp, err := filepath.Abs(fp)
			if err != nil {
				log.Println(err)
				continue
			}
			fps, err := readPathFiles(rfp, FlagRecursive)
			if err != nil {
				log.Println(err)
				continue
			}
			for _, f := range fps {
				select {
				case <-ctx.Done():
					return
				case ch <- f:
				}
			}
		}
	}()
	return ch
}

func readPathFiles(p string, recurse bool) ([]string, error) {
	fi, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return []string{p}, nil
	}

	if !FlagRecursive {
		return nil, fmt.Errorf("%s is a directory", p)
	}

	fis, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}
	var fps []string
	var dfps []string
	for _, dfi := range fis {
		if !dfi.IsDir() {
			fps = append(fps, path.Join(p, dfi.Name()))
			continue
		}
		f, err := readPathFiles(path.Join(p, dfi.Name()), recurse)
		if err != nil {
			return nil, err
		}
		dfps = append(dfps, f...)
	}
	return append(fps, dfps...), nil
}
