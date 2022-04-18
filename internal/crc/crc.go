package crc

import (
	"errors"
	"hash/crc32"
	"io"
)

type Crc32Hasher struct {
	table *crc32.Table
}

func NewCrc32Hasher() *Crc32Hasher {
	return &Crc32Hasher{
		table: crc32.MakeTable(crc32.Castagnoli),
	}
}

func (h *Crc32Hasher) Hash(r io.Reader) (uint32, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, errors.New("unable to hash a size of 0 data")
	}

	return crc32.Checksum(data, h.table), nil
}
