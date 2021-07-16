package format

import (
	gofmt "go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Format format go source file in place
func Format(fpath string) error {

	in, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}

	out, err := gofmt.Source(in)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fpath, out, 0644)
}

// FormatDir format go source directory in place
func FormatDir(dir string) error {
	err := filepath.Walk(dir, func(fpath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(fpath, ".go") && !info.IsDir() {
			return Format(fpath)
		}
		return nil
	})

	return err
}
