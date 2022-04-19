package fdup

import (
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/enkeyz/go-fdup/internal/hash"
)

func TestSearch(t *testing.T) {
	data := "Hello, World"
	hasher := hash.NewCrc32Hasher()
	hash, _ := hasher.Hash(strings.NewReader(data))

	t.Run("searching for duplicate files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello.txt": {
				Data: []byte(data),
			},
			filepath.Join("subdir", "hello_copy.txt"): {
				Data: []byte(data),
			},
		}

		fdup := NewFdup(fs)
		got, err := fdup.Search()

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(got[hash]) != 2 {
			t.Errorf("got %d files, but expected %d", len(got[hash]), 2)
		}
	})
}
