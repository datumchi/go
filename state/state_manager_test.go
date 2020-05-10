package state_test

import (
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/state"
	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateManager", func() {


	Describe("Creating a StateManager for CollaborationState", func() {

		Context("Using an empty Collaboration State", func() {

			collabState := protocol.CollaborationState{}
			collabState.RootChildren = make(map[string][]byte)
			stateManager, err := state.CreateStateManager(collabState)

			It("Reads no entries", func() {

				Expect(err).To(BeNil())
				Expect(stateManager.TotalRootEntries()).To(Equal(0))

			})

		})

		Context("Using a few root entries with simple values, no permissions", func() {

			collabState := GenerateSimpleStateNodesCollabState()

			stateManager, err := state.CreateStateManager(collabState)

			It("Reads three entries", func() {

				Expect(err).To(BeNil())
				Expect(stateManager.TotalRootEntries()).To(Equal(3))


			})


			It("Had valid values", func() {

				oneValue := stateManager.GetValueForPath([]string {"one"})
				twoValue := stateManager.GetValueForPath([]string {"two"})
				threeValue := stateManager.GetValueForPath([]string {"three"})

				Expect(oneValue).To(Equal([]byte("!!ONE!!")))
				Expect(twoValue).To(Equal([]byte("!!TWO!!")))
				Expect(threeValue).To(Equal([]byte("!!THREE!!")))

			})




		})

		Context("Using a few root entries with children, no permissions", func() {

			collabState := GenerateMultilevelStateNodesCollabStateNoPermissions()
			stateManager, err := state.CreateStateManager(collabState)

			It("Reads three entries", func() {
				Expect(err).To(BeNil())
				Expect(stateManager.TotalRootEntries()).To(Equal(3))
			})

			It("Is able to get a simple single-level value", func() {

				oneValue := stateManager.GetValueForPath([]string {"one"})
				Expect(oneValue).To(Equal([]byte("!!ONE!!")))

			})

			It("Is able to get a lower-level leaf value", func() {

				oneOneTwoValue := stateManager.GetValueForPath([]string {"one", "oneone", "oneonetwo"})
				Expect(oneOneTwoValue).ToNot(BeNil())
				Expect(oneOneTwoValue).To(Equal([]byte("!!ONEONETWO!!")))

			})

			It("Returns a nil for an unknown value", func() {

				oneOneTwoValue := stateManager.GetValueForPath([]string {"one", "one1one", "oneone5two"})
				Expect(oneOneTwoValue).To(BeNil())

			})



		})




	})

	Describe("Generating a hash for the state", func() {

		Context("A simple single-level set of values", func() {

			collabState := GenerateSimpleStateNodesCollabState()
			encodedState, _ := proto.Marshal(&collabState)
			hashState := sha512.HashData(encodedState)

			stateManager, err := state.CreateStateManager(collabState)

			PIt("Should generate a correct hash for the state", func() {
				Expect(err).To(BeNil())
				Expect(stateManager.Hash()).To(Equal(hashState))
			})

		})





	})


})


func GenerateSimpleStateNodesCollabState() protocol.CollaborationState {

	collabState := protocol.CollaborationState{}
	collabState.RootChildren = make(map[string][]byte)

	nodeOne := protocol.CollaborationState_Node{
		Label: "one",
		Value: []byte("!!ONE!!"),
	}
	nodeOneEncoded, _ := proto.Marshal(&nodeOne)

	nodeTwo := protocol.CollaborationState_Node{
		Label: "two",
		Value: []byte("!!TWO!!"),
	}
	nodeTwoEncoded, _ := proto.Marshal(&nodeTwo)

	nodeThree := protocol.CollaborationState_Node{
		Label: "three",
		Value: []byte("!!THREE!!"),
	}
	nodeThreeEncoded, _ := proto.Marshal(&nodeThree)

	collabState.RootChildren["one"] = nodeOneEncoded
	collabState.RootChildren["two"] = nodeTwoEncoded
	collabState.RootChildren["three"] = nodeThreeEncoded

	return collabState

}

func GenerateMultilevelStateNodesCollabStateNoPermissions() protocol.CollaborationState {

	collabState := protocol.CollaborationState{}
	collabState.RootChildren = make(map[string][]byte)


	nodeOneOneTwo := protocol.CollaborationState_Node{
		Label: "oneonetwo",
		Value: []byte("!!ONEONETWO!!"),
	}
	nodeOneOneTwoEncoded, _ := proto.Marshal(&nodeOneOneTwo)

	nodeOneOne := protocol.CollaborationState_Node{
		Label: "oneone",
		Children: make(map[string][]byte),
	}
	nodeOneOne.Children["oneonetwo"] = nodeOneOneTwoEncoded
	nodeOneOneEncoded, _ := proto.Marshal(&nodeOneOne)

	nodeOne := protocol.CollaborationState_Node{
		Label: "one",
		Value: []byte("!!ONE!!"),
		Children: make(map[string][]byte),
	}
	nodeOne.Children["oneone"] = nodeOneOneEncoded
	nodeOneEncoded, _ := proto.Marshal(&nodeOne)

	nodeTwo := protocol.CollaborationState_Node{
		Label: "two",
		Value: []byte("!!TWO!!"),
	}
	nodeTwoEncoded, _ := proto.Marshal(&nodeTwo)

	nodeThree := protocol.CollaborationState_Node{
		Label: "three",
		Value: []byte("!!THREE!!"),
	}
	nodeThreeEncoded, _ := proto.Marshal(&nodeThree)

	collabState.RootChildren["one"] = nodeOneEncoded
	collabState.RootChildren["two"] = nodeTwoEncoded
	collabState.RootChildren["three"] = nodeThreeEncoded

	return collabState

}

