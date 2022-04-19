package hash

import (
	"bytes"
	"fmt"
	"testing"
	"testing/fstest"
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
		fs := fstest.MapFS{
			"hello.txt": {
				Data: []byte(data),
			},
		}

		testFile, err := fs.Open("hello.txt")
		if err != nil {
			t.Fatal("error opening file")
		}

		got, _ := hasher.Hash(testFile)
		var expected uint32 = 1080205678

		if got != expected {
			t.Errorf("got %d, but expected %d", got, expected)
		}
	})
}
