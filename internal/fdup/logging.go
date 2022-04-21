package fdup

import "fmt"

func PrintResult(hashedFileMap HashedFileMap) {
	for _, fileInfoSlice := range hashedFileMap {
		if len(fileInfoSlice) == 1 {
			continue
		}

		fmt.Println("Duplicate(s) found!")
		for _, fileInfo := range fileInfoSlice {
			fmt.Printf("\thash: %d, file: %s\n", fileInfo.Hash, fileInfo.Path)
		}
	}
}
