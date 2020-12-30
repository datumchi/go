package createcollab

import (
	"github.com/datumchi/go/core/context"
	"github.com/datumchi/go/generated/protocol"
)

func DoCreateCollaboration(nodeContext context.NodeContext, collab protocol.Collaboration) {

	// Determine whether this is a virgin, NeedReturn, or IsReturn
	if len(collab.CollaborationOriginationSignatures) == 1 {

	}

}
