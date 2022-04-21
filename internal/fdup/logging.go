package fdup

import (
	"fmt"
	"io"
)

func PrintResult(w io.Writer, hashedFileMap HashedFileMap) {
	for hash, fileInfoSlice := range hashedFileMap {
		if len(fileInfoSlice) == 1 {
			continue
		}

		fmt.Fprintf(w, "Hash: %d\n", hash)
		for _, fileInfo := range fileInfoSlice {
			fmt.Fprintf(w, "\tfile: %s\n", fileInfo.Path)
		}
	}
}
