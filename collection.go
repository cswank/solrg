package solrg

import "github.com/cswank/solrg/internal/zk"

type Collection struct {
	collection zk.Collection
}

func NewCollection(name string) (*Collection, error) {
	c, err := zk.GetCollection(name)
	return &Collection{collection: c}, err
}

func (c *Collection) Leader(id string) (string, error) {

}
