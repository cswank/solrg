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
				Shards: map[string]zk.Shard{
					"shard1": {
						Range: "75550000-7fffffff",
						State: "active",
						Replicas: map[string]zk.Replica{
							"core_node27": {
								BaseURL: "http://172.19.0.3:8983/solr",
								State:   "active",
								Leader:  "true",
							},
							"core_node33": {
								BaseURL: "http://172.19.0.2:8983/solr",
								State:   "active",
							},
						},
					},
					"shard2": {
						Range: "75550000-7fffffff",
						State: "active",
						Replicas: map[string]zk.Replica{
							"core_node1": {
								BaseURL: "http://172.19.0.1:8983/solr",
								State:   "active",
								Leader:  "true",
							},
							"core_node3": {
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
