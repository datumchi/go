package state

import (
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/generated/protocol"
	"github.com/gogo/protobuf/proto"
)

type StateManager struct {
	CollabState protocol.CollaborationState
}

func CreateStateManager(collaborationState protocol.CollaborationState) (StateManager, error) {

	sm := StateManager{
		CollabState: collaborationState,
	}

	return sm, nil

}

func (sm StateManager) Hash() []byte {

	st := sm.CollabState
	encodedState, err := proto.Marshal(&st)
	if err != nil {
		return nil
	}

	hashState := sha512.HashData(encodedState)
	return hashState

}

func (sm StateManager) TotalRootEntries() int {

	keys := make([]string, 0, len(sm.CollabState.RootChildren))
	for k := range sm.CollabState.RootChildren {
		keys = append(keys, k)
	}

	return len(keys)

}

func (sm StateManager) GetValueForPath(paths []string) []byte {

	if len(paths) == 0 {
		return nil
	}

	// Get the first node in sequence
	var stateNode protocol.CollaborationState_Node
	if sm.CollabState.RootChildren[paths[0]] != nil {

		stateNodeEncoded := sm.CollabState.RootChildren[paths[0]]
		err := proto.Unmarshal(stateNodeEncoded, &stateNode)
		if err != nil {
			return nil
		}

	}

	if len(paths) == 1 {
		return stateNode.Value
	}

	// dive into structure
	for _,pathName := range paths[1:] {

		encodedChild := stateNode.Children[pathName]
		if encodedChild == nil {
			return nil
		}
		err := proto.Unmarshal(encodedChild, &stateNode)
		if err != nil {
			return nil
		}

	}

	return stateNode.Value

}


