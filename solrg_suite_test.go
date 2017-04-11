package solrg_test

import (
	"github.com/cswank/solrg/internal/zk"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSolrg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solrg Suite")
}

var _ = Describe("solrg", func() {
	Describe("GetShard", func() {
		var (
			col zk.Collection
		)

		BeforeEach(func() {
			col = zk.Collection{
				Leaders: map[string]string{
					"shard1": "http://172.19.0.2:8983/solr",
					"shard2": "http://172.19.0.6:8983/solr",
					"shard3": "http://172.19.0.2:8983/solr",
				},
				Shards: map[string]zk.Shard{
					"shard1": {
						Range: "80000000-d554ffff",
						State: "active",
						Replicas: map[string]zk.Replica{
							"core_node2": {
								BaseURL: "http://172.19.0.2:8983/solr",
								State:   "active",
								Leader:  "true",
							},
							"core_node6": {
								BaseURL: "http://172.19.0.4:8983/solr",
								State:   "active",
							},
						},
					},
					"shard2": {
						Range: "d5550000-2aa9ffff",
						State: "active",
						Replicas: map[string]zk.Replica{
							"core_node1": {
								BaseURL: "http://172.19.0.6:8983/solr",
								State:   "active",
								Leader:  "true",
							},
							"core_node4": {
								BaseURL: "http://172.19.0.3:8983/solr",
								State:   "active",
							},
						},
					},
					"shard3": {
						Range: "2aaa0000-7fffffff",
						State: "active",
						Replicas: map[string]zk.Replica{
							"core_node3": {
								BaseURL: "http://172.19.0.2:8983/solr",
								State:   "active",
								Leader:  "true",
							},
							"core_node5": {
								BaseURL: "http://172.19.0.4:8983/solr",
								State:   "active",
							},
						},
					},
				},
			}
		})
	})
})
