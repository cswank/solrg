package zk

var (
	GetCollection func(string) (Collection, error)
)

type Replica struct {
	Core     string `json:"core"`
	BaseURL  string `json:"base_url"`
	NodeName string `json:"node_name"`
	State    string `json:"state"`
	Leader   string `json:"leader"`
}

type Shard struct {
	Range    string             `json:"range"`
	State    string             `json:"state"`
	Replicas map[string]Replica `json:"replicas"`
}

type Collection struct {
	ReplicationFactor string           `json:"replicationFactor"`
	MaxShardsPerNode  string           `json:"maxShardsPerNode"`
	Shards            map[string]Shard `json:"shards"`
	Router            struct {
		Name string `json:"name"`
	} `json:"router"`

	//Leaders   shard   host
	Leaders map[string]string
}
