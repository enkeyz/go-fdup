package fdup

import (
	"errors"
	"io/fs"

	"github.com/enkeyz/go-fdup/internal/hash"
)

type FileInfo struct {
	Name     string
	FullPath string
	Size     int64
}

type HashedFileMap map[uint32][]FileInfo

type Fdup struct {
	f      fs.FS
	hasher *hash.Crc32Hasher
}

// creating a new instance of fdup to use to find duplicate files
func NewFdup(f fs.FS) *Fdup {
	return &Fdup{
		f,
		hash.NewCrc32Hasher(),
	}
}

// the main method to search for files
func (fd *Fdup) Search() (HashedFileMap, error) {
	fileInfos, err := fd.getAllFileInfo()
	if err != nil {
		return nil, err
	}

	if len(fileInfos) == 0 {
		return nil, errors.New("no file found")
	}

	fmap, err := fd.search(fileInfos)
	if err != nil {
		return nil, err
	}

	if !fd.duplicatesExists(fmap) {
		return nil, errors.New("no duplicate files found")
	}

	return fmap, nil
}

func (fd *Fdup) search(fileInfos []FileInfo) (HashedFileMap, error) {
	fsmap := make(map[int64][]FileInfo)
	for _, fileInfo := range fileInfos {
		fsmap[fileInfo.Size] = append(fsmap[fileInfo.Size], fileInfo)
	}

	hfmap := make(map[uint32][]FileInfo)
	for _, fileInfos := range fsmap {
		if len(fileInfos) <= 1 {
			continue
		}

		for _, fileInfo := range fileInfos {
			file, err := fd.f.Open(fileInfo.FullPath)
			if err != nil {
				return nil, err
			}

			// TODO only hash when the first x bytes are equal
			hash, err := fd.hasher.Hash(file)
			if err != nil {
				continue
			}

			hfmap[hash] = append(hfmap[hash], fileInfo)
			file.Close()
		}
	}

	return hfmap, nil
}

// check if there are more then one file with the same hash in the map
func (fd *Fdup) duplicatesExists(hashedFileMap HashedFileMap) bool {
	for _, slice := range hashedFileMap {
		if len(slice) > 1 {
			return true
		}
	}

	return false
}

// get all files in the given directory by the user
func (fd *Fdup) getAllFileInfo() ([]FileInfo, error) {
	files := make([]FileInfo, 0)

	err := fs.WalkDir(fd.f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || d.Type() == fs.ModeSymlink {
			return nil
		}

		finfo, err := d.Info()
		if err != nil {
			return err
		}

		files = append(files, FileInfo{Name: finfo.Name(), FullPath: path, Size: finfo.Size()})

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, err
}
