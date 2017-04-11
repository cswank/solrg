package key_test

import (
	"github.com/cswank/solrg/internal/key"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKey(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Key Suite")
}

var _ = Describe("Key", func() {
	DescribeTable("hashes", func(id string, expected int32) {
		k, err := key.New(id)
		Expect(err).To(BeNil())
		Expect(k.Hash()).To(Equal(expected))
	},
		Entry("", "a!b", int32(1009090051)),
		Entry("", "list_segmentation_76661!me@hi.com", int32(1617920828)),
		Entry("", "76661!craig@localhost", int32(1815961548)),
	)
})
