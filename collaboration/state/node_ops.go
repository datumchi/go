package state

import (
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/utility/logger"
	"github.com/golang/protobuf/proto"
	"strings"
)

func GetOrCreateNodePath(rootNode *protocol.CollaborationState_Node, nodePath string) protocol.CollaborationState_Node {

	if rootNode == nil {
		return protocol.CollaborationState_Node{}
	}

	if nodePath == "" || nodePath == "/" {
		return *rootNode
	}

	nodePathElements := strings.Split(nodePath, "/")
	var currentNode = *rootNode
	for _,nodePathElement := range nodePathElements {

		// Root (this one) is first, skip that.
		// Get next item from list
		var node protocol.CollaborationState_Node
		nodeBytes := currentNode.Children[nodePathElement]
		if len(nodeBytes) == 0 {
			// create the children!
			//childNode := protocol.CollaborationState_Node{Label: nodePathElement}
			//rootNode.Children[nodePathElement] =
		} else {

			err := proto.Unmarshal(nodeBytes, &node)
			if err != nil {
				return protocol.CollaborationState_Node{}
			}

			


		}

		err := proto.Unmarshal(nodeBytes, &currentNode)

	}


	return protocol.CollaborationState_Node{}

}
