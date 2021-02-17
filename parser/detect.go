package parser

import (
	"bufio"
	"bytes"
	"github.com/eurozulu/4mat/formats"
	"io/ioutil"
	"path"
	"strings"
)

var DetectContent bool

func DetectFormat(p string) (formats.Format, error) {
	if DetectContent {
		return detectByContent(p)
	}
	return detectByFileName(p), nil
}

func detectByFileName(p string) formats.Format {
	e := path.Ext(p)
	if e == "" {
		return formats.FormatUnknown
	}
	var fm formats.Format
	if err := fm.UnmarshalText([]byte(strings.ToLower(e))); err != nil {
		return formats.FormatUnknown
	}
	return fm
}

func detectByContent(p string) (formats.Format, error) {
	by, err := ioutil.ReadFile(p)
	if err != nil {
		return formats.FormatUnknown, err
	}
	scn := bufio.NewScanner(bytes.NewReader(by))
	for scn.Scan() {
		panic("Not yet implemented")
	}
	return formats.FormatUnknown, nil
}
