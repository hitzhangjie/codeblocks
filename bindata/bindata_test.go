package main

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/hitzhangjie/codeblocks/tar"
	"github.com/stretchr/testify/assert"
)

func Test_readFromInputSource(t *testing.T) {
	t.Run("case invalid input", func(t *testing.T) {
		p := gomonkey.NewPatches()
		p.ApplyFunc(os.Lstat, func(name string) (os.FileInfo, error) {
			return nil, errors.New("fake error")
		})
		defer p.Reset()
		data, err := readFromInputSource("")
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})

	t.Run("case tar error", func(t *testing.T) {
		p := gomonkey.NewPatches()
		p.ApplyFunc(os.Lstat, func(name string) (os.FileInfo, error) {
			return nil, nil
		})
		p.ApplyFunc(tar.Tar, func(src string, writers ...io.Writer) error {
			return errors.New("fake error")
		})
		defer p.Reset()
		data, err := readFromInputSource("")
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})

	t.Run("case success", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		_, err = readFromInputSource(dir)
		assert.Nil(t, err)
	})
}
