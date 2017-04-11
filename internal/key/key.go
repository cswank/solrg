package key

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

type Key struct {
	key    string
	hashes []int32
	masks  []int32
	bits   []uint32
}

func New(key string) (*Key, error) {
	parts := strings.Split(key, "!")
	if len(parts) != 2 {
		return nil, fmt.Errorf("error: could not parse %s", key)
	}
	return &Key{
		key: key,
		hashes: []int32{
			int32(murmur3.Sum32([]byte(parts[0]))),
			int32(murmur3.Sum32([]byte(parts[1]))),
		},
		masks: []int32{
			(-1 << (32 - 16)),
			65535,
		},
	}, nil
}

func (k *Key) Hash() int32 {
	return (k.hashes[0] & k.masks[0]) | (k.hashes[1] & k.masks[1])
}
