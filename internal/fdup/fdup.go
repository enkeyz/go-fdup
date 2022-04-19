package fdup

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/enkeyz/go-fdup/internal/hash"
)

type HashedFileInfo struct {
	Name string
	Path string
	Size int64
	Hash uint32
}

type HashedFileMap map[uint32][]*HashedFileInfo

type Fdup struct {
	f fs.FS
}

func NewFdup(f fs.FS) *Fdup {
	return &Fdup{
		f,
	}
}

// TODO only hash files if they're the same size
func (fd *Fdup) Search() (HashedFileMap, error) {
	filePaths, err := fd.getAllFiles()
	if err != nil {
		return nil, err
	}

	hasher := hash.NewCrc32Hasher()
	fmap := make(map[uint32][]*HashedFileInfo)
	for _, filePath := range filePaths {
		file, err := fd.f.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		hash, err := hasher.Hash(file)
		if err != nil {
			return nil, err
		}

		fInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}

		fmap[hash] = append(fmap[hash], &HashedFileInfo{fInfo.Name(), filePath, fInfo.Size(), hash})
	}

	return fmap, nil
}

func (fd *Fdup) PrintDuplicateFiles(hashedFileMap HashedFileMap) {
	for _, fileInfoSlice := range hashedFileMap {
		if len(fileInfoSlice) == 1 {
			continue
		}

		fmt.Println("Duplicate(s) found!")
		for _, fileInfo := range fileInfoSlice {
			fmt.Printf("	hash: %d, file: %s, size: %d bytes\n", fileInfo.Hash, fileInfo.Path, fileInfo.Size)
		}
	}
}

func (fd *Fdup) getAllFiles() ([]string, error) {
	files := []string{}

	fs.WalkDir(fd.f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		files = append(files, filepath.Join(".", path))

		return nil
	})

	return files, nil
}
