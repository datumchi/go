package state_test

import (
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/state"
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"


)

var _ = Describe("StateHashing", func() {


	Describe("Hashing the State of a Collaboration", func() {

		var collaborationState protocol.CollaborationState

		BeforeEach(func() {

			collaborationState = protocol.CollaborationState{
				RootChildren: make(map[string][]byte),
			}

		})

		It("Should be repeatable", func() {

			nodeOne := protocol.CollaborationState_Node{
				Label: "data",
				Value: []byte("data-value-one"),
			}
			nodeOneRaw, _ := proto.Marshal(&nodeOne)
			
			nodeTwo := protocol.CollaborationState_Node{
				Label: "dataAgain",
				Value: []byte("data-value-two"),
			}
			nodeTwoRaw, _ := proto.Marshal(&nodeTwo)

			collaborationState.RootChildren["data"] = nodeOneRaw
			collaborationState.RootChildren["dataAgain"] = nodeTwoRaw

			collabHashOne, err := state.HashCollaborationState(&collaborationState)
			Expect(err).To(BeNil())
			Expect(collabHashOne).ToNot(BeNil())


			nodeOneClone := protocol.CollaborationState_Node{
				Label: "data",
				Value: []byte("data-value-one"),
			}
			nodeOneCloneRaw, _ := proto.Marshal(&nodeOneClone)

			nodeTwoClone := protocol.CollaborationState_Node{
				Label: "dataAgain",
				Value: []byte("data-value-two"),
			}
			nodeTwoCloneRaw, _ := proto.Marshal(&nodeTwoClone)

			collaborationStateClone := protocol.CollaborationState{
				RootChildren: make(map[string][]byte),
			}
			collaborationStateClone.RootChildren["data"] = nodeOneCloneRaw
			collaborationStateClone.RootChildren["dataAgain"] = nodeTwoCloneRaw

			collabCloneHashOne, err := state.HashCollaborationState(&collaborationStateClone)
			Expect(err).To(BeNil())
			Expect(collabCloneHashOne).ToNot(BeEmpty())
			Expect(collabCloneHashOne).To(Equal(collabHashOne))

		})




	})



})
