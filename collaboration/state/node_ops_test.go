package state_test

import (
	"github.com/datumchi/go/collaboration/state"
	"github.com/datumchi/go/generated/protocol"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NodeOps", func() {

	Describe("Creating a node where it doesn't yet exist", func() {

		var rootNode = protocol.CollaborationState_Node{Label: ""}

		It("Lets you create a node if you GetOrCreate and the path doesn't yet exist", func() {

			createdNode := state.GetOrCreateNodePath(&rootNode, "/this/is/my/test/path")
			Expect(createdNode.Label).To(Equal("path"))


		})

	})


})
