package solrg

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
	return "", nil
}
