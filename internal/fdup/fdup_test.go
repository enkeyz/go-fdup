package fdup

import (
	"testing"
	"testing/fstest"
)

func TestSearch(t *testing.T) {
	data := "Hello, World"
	var hash uint32 = 1080205678

	t.Run("searching for duplicate files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello.txt": {
				Data: []byte(data),
			},
			"subdir/hello_copy.txt": {
				Data: []byte(data),
			},
		}

		fdup := NewFdup(fs)
		got, _ := fdup.Search()

		if len(got[hash]) != 2 {
			t.Errorf("got %d files, but expected %d", len(got[hash]), 2)
		}
	})
}
