package crc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCrc32Hash(t *testing.T) {
	data := "Hello, World"
	hasher := NewCrc32Hasher()

	t.Run("calculate hash of bytes.Buffer", func(t *testing.T) {
		buff := &bytes.Buffer{}
		fmt.Fprint(buff, data)

		got, _ := hasher.Hash(buff)
		var expected uint32 = 1080205678

		if got != expected {
			t.Errorf("got %d, but expected %d", got, expected)
		}
	})

	t.Run("hashing of an empty sized bytes.Buffer", func(t *testing.T) {
		buff := &bytes.Buffer{}

		_, err := hasher.Hash(buff)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("hashing contents of a file", func(t *testing.T) {
		tmpFile, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatal("error creating temp file")
		}
		defer os.Remove(tmpFile.Name())

		fmt.Fprint(tmpFile, data)
		tmpFile.Seek(0, io.SeekStart)

		got, _ := hasher.Hash(tmpFile)
		var expected uint32 = 1080205678

		if got != expected {
			t.Errorf("got %d, but expected %d", got, expected)
		}

		if err := tmpFile.Close(); err != nil {
			t.Fatal("error closing temp file")
		}
	})
}
