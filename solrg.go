package solrg

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

type Solrg struct {
	shards int
}

func New(shards int) *Solrg {
	return &Solrg{
		shards: shards,
	}
}

//Host uses the shard key on the id to find the leader
func (s *Solrg) Host(id string) (string, error) {
	parts := strings.Split(id, "!")
	if len(parts) != 2 {
		return "", fmt.Errorf("no shard key found in id: %s", id)
	}

	routeHash := murmur3.Sum32([]byte(parts[0]))
	idHash := murmur3.Sum32([]byte(parts[1]))

	return "", nil
}
