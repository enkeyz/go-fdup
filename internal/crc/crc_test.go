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
	t.Run("calculate hash of bytes.Buffer", func(t *testing.T) {
		buff := &bytes.Buffer{}
		fmt.Fprint(buff, "Hello, World")

		got, _ := Crc32Hash(buff)
		var expected uint32 = 1080205678

		if got != expected {
			t.Errorf("got %d, but expected %d", got, expected)
		}
	})

	t.Run("hashing of an empty sized bytes.Buffer", func(t *testing.T) {
		buff := &bytes.Buffer{}

		_, err := Crc32Hash(buff)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("hashing contents of a file", func(t *testing.T) {
		data := []byte("Hello, World")
		tmpFile, err := ioutil.TempFile("", "testfile")
		if err != nil {
			t.Fatal("error creating temp file")
		}
		defer os.Remove(tmpFile.Name())
		if _, err := tmpFile.Write(data); err != nil {
			t.Fatal("error writing into temp file")
		}
		tmpFile.Seek(0, io.SeekStart)

		got, _ := Crc32Hash(tmpFile)
		var expected uint32 = 1080205678

		if got != expected {
			t.Errorf("got %d, but expected %d", got, expected)
		}

		if err := tmpFile.Close(); err != nil {
			t.Fatal("error closing temp file")
		}
	})
}
