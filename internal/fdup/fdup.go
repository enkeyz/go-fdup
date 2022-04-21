package fdup

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/enkeyz/go-fdup/internal/hash"
)

type HashedFileInfo struct {
	Path string
	Hash uint32
}

type HashedFileMap map[uint32][]HashedFileInfo

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

func (fd *Fdup) Search() (HashedFileMap, error) {
	filePaths, err := fd.getAllFilePath()
	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, errors.New("no file found")
	}

	fmap, err := fd.search(filePaths)
	if err != nil {
		return nil, err
	}

	if !fd.duplicatesExists(fmap) {
		return nil, errors.New("no duplicate files found")
	}

	return fmap, nil
}

func (fd *Fdup) search(filePaths []string) (HashedFileMap, error) {
	fsmap := make(map[int64][]string)
	for _, filePath := range filePaths {
		file, err := fd.f.Open(filePath)
		if err != nil {
			return nil, err
		}

		fInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}

		fsmap[fInfo.Size()] = append(fsmap[fInfo.Size()], filePath)
		file.Close()
	}

	hfmap := make(map[uint32][]HashedFileInfo)
	for _, files := range fsmap {
		if len(files) <= 1 {
			continue
		}

		for _, filePath := range files {
			file, err := fd.f.Open(filePath)
			if err != nil {
				return nil, err
			}

			hash, err := fd.hasher.Hash(file)
			if err != nil {
				continue
			}

			hfmap[hash] = append(hfmap[hash], HashedFileInfo{Path: filePath, Hash: hash})
			file.Close()
		}
	}

	return hfmap, nil
}

// check if there are more then one file with the same hash in the whole map
func (fd *Fdup) duplicatesExists(hashedFileMap HashedFileMap) bool {
	for _, slice := range hashedFileMap {
		if len(slice) > 1 {
			return true
		}
	}

	return false
}

func (fd *Fdup) getAllFilePath() ([]string, error) {
	files := []string{}

	fs.WalkDir(fd.f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		files = append(files, filepath.Join(".", path))

		return nil
	})

	return files, nil
}
