package state

import (
	"github.com/datumchi/go/crypto/sha256"
	"github.com/datumchi/go/generated/protocol"
)

func HashCollaborationState(state *protocol.CollaborationState) ([]byte, error) {

	var content []byte
	for _,v := range state.RootChildren {
		content = append(content, v...)
	}

	hashedData := sha256.HashData(content)
	return hashedData, nil



}



