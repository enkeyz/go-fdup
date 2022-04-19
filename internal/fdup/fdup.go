package fdup

import (
	"errors"
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
	f      fs.FS
	hasher *hash.Crc32Hasher
}

func NewFdup(f fs.FS) *Fdup {
	return &Fdup{
		f,
		hash.NewCrc32Hasher(),
	}
}

// TODO only hash files if they're the same size
func (fd *Fdup) Search() (HashedFileMap, error) {
	filePaths, err := fd.getAllFilePath()
	if err != nil {
		return nil, err
	}

	fmap, err := fd.search(filePaths)
	if err != nil {
		return nil, err
	}

	if fd.checkForDuplicates(fmap) {
		return nil, errors.New("no duplicate files found")
	}

	return fmap, nil
}

func (fd *Fdup) search(filePaths []string) (HashedFileMap, error) {
	fmap := make(map[uint32][]*HashedFileInfo)
	for _, filePath := range filePaths {
		file, err := fd.f.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		hash, err := fd.hasher.Hash(file)
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

// check if there are more then one file with the same hash in the whole map
func (fd *Fdup) checkForDuplicates(hashedFileMap HashedFileMap) bool {
	for _, slice := range hashedFileMap {
		if len(slice) > 1 {
			return false
		}
	}

	return true
}

func (fd *Fdup) getAllFilePath() ([]string, error) {
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
