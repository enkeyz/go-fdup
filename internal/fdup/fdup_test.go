package fdup

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestSearch(t *testing.T) {
	data := "Hello, World"

	fs := fstest.MapFS{
		"hello.txt": {
			Data: []byte(data),
		},
		"subdir/bye.txt": {
			Data: []byte(data),
		},
	}
	fdup := NewFdup(fs)

	got, _ := fdup.Search()
	expected := HashedFileMap{
		1080205678: []*HashedFileInfo{
			{Name: "hello.txt", Path: "hello.txt", Size: int64(len(data)), Hash: 1080205678},
			{Name: "bye.txt", Path: "subdir/bye.txt", Size: int64(len(data)), Hash: 1080205678},
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, but expected %v", got, expected)
	}
}
