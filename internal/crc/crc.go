package crc

import (
	"errors"
	"hash/crc32"
	"io"
)

func Crc32Hash(r io.Reader) (uint32, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, errors.New("unable to hash a size of 0 data")
	}

	crc32q := crc32.MakeTable(crc32.Castagnoli)
	return crc32.Checksum(data, crc32q), nil
}
