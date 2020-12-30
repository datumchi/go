package state_test

import (
	"github.com/datumchi/go/collaboration/state"
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/generated/protocol"
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CalculateCollaborationStateHash", func() {

	Describe("Calculating the hash of a single-level node", func() {

		// single level node:
		// root (/) -> A(label="one", value="1"), B(label="two", value="2"), C(label="three", value="3")
		var singleLevelNode protocol.CollaborationState_Node
		singleLevelNode.Children = make(map[string][]byte)

		BeforeEach(func() {

			// A Node
			aNode := protocol.CollaborationState_Node{
				Label: "one",
				Value: []byte("1"),
			}
			aNodeBytes, _ := proto.Marshal(&aNode)
			singleLevelNode.Children["one"] = aNodeBytes

			// B Node
			bNode := protocol.CollaborationState_Node{
				Label: "two",
				Value: []byte("2"),
			}
			bNodeBytes, _ := proto.Marshal(&bNode)
			singleLevelNode.Children["two"] = bNodeBytes

			// C Node
			cNode := protocol.CollaborationState_Node{
				Label: "three",
				Value: []byte("3"),
			}
			cNodeBytes, _ := proto.Marshal(&cNode)
			singleLevelNode.Children["three"] = cNodeBytes

		})

		It("Should create the expected hash", func() {

			// What is expected hash?
			var data []byte
			data = append(data, sha512.HashData([]byte("/"))...)

			dataOne := sha512.HashData([]byte("one"))
			dataOne = append(dataOne, sha512.HashData([]byte("1"))...)
			dataOneHash := sha512.HashData(dataOne)
			data = append(data, dataOneHash...)

			dataThree := sha512.HashData([]byte("three"))
			dataThree = append(dataThree, sha512.HashData([]byte("3"))...)
			dataThreeHash := sha512.HashData(dataThree)
			data = append(data, dataThreeHash...)

			dataTwo := sha512.HashData([]byte("two"))
			dataTwo = append(dataTwo, sha512.HashData([]byte("2"))...)
			dataTwoHash := sha512.HashData(dataTwo)
			data = append(data, dataTwoHash...)

			expectedHash := sha512.HashData(data)
			calculatedHash := state.CalculateCollaborationStateHash(singleLevelNode)

			Expect(calculatedHash).To(Equal(expectedHash))

		})

	})

})
