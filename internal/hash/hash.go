package hash

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

func Hash(id string) (int32, error) {
	parts := strings.Split(id, "!")
	if len(parts) != 2 {
		return 0, fmt.Errorf("error: could not parse %s", id)
	}

	hashes := []int32{
		int32(murmur3.Sum32([]byte(parts[0]))),
		int32(murmur3.Sum32([]byte(parts[1]))),
	}
	masks := []int32{
		(-1 << (32 - 16)), // -10000000000000000
		65535,             // 1111111111111111
	}

	return (hashes[0] & masks[0]) | (hashes[1] & masks[1]), nil
}
