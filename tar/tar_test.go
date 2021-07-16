package tar

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTar(t *testing.T) {
	// setup
	tmp := os.TempDir()
	dir := filepath.Join(tmp, "bindata")
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fp := filepath.Join(dir, "test.txt")
	txt := []byte("helloworld")
	err = ioutil.WriteFile(fp, txt, 0666)
	if err != nil {
		panic(err)
	}

	// tar
	buf := bytes.Buffer{}
	err = Tar(fp, &buf)
	assert.Nil(t, err)
	assert.NotZero(t, buf.Len())

	// untar
	reader := bytes.NewReader(buf.Bytes())
	dst := filepath.Join(dir, "test2.txt")
	err = Untar(dst, reader)
	assert.Nil(t, err)

	dat, err := ioutil.ReadFile(dst)
	assert.Nil(t, err)
	assert.Equal(t, txt, dat)

	os.RemoveAll(dir)
}

func TestTar_Install(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	src := wd

	// tar
	buf := bytes.Buffer{}
	err = Tar(src, &buf)
	assert.Nil(t, err)
	assert.NotZero(t, buf.Len())

	// untar
	tmp := os.TempDir()
	dst := filepath.Join(tmp, "bindata")
	defer os.RemoveAll(dst)

	err = Untar(dst, &buf)
	assert.Nil(t, err)

	// compare file list
	srcFileSet := map[string]struct{}{}
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		path = strings.TrimPrefix(path, src)
		srcFileSet[path] = struct{}{}
		return nil
	})
	dstFileSet := map[string]struct{}{}
	filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {
		path = strings.TrimPrefix(path, dst)
		dstFileSet[path] = struct{}{}
		return nil
	})
	assert.Equal(t, srcFileSet, dstFileSet)
}
