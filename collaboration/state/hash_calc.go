package state

import (
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/generated/protocol"
	"github.com/golang/protobuf/proto"
	"sort"
)

var (
	ROOT_LABEL_HASH = sha512.HashData([]byte("/"))
)

func CalculateCollaborationStateHash(stateNode protocol.CollaborationState_Node) []byte {

	// Holds the data to be hashed eventually
	var nodeHashInputData []byte

	// Go ahead and hash the node's label and value and append to data
	// If we are root (label = ""), append hash of "/"
	if stateNode.Label == "" {

		nodeHashInputData = append(nodeHashInputData, ROOT_LABEL_HASH...)

	} else {

		// Label
		hashedNodeLabel := sha512.HashData([]byte(stateNode.Label))
		nodeHashInputData = append(nodeHashInputData, hashedNodeLabel...)

		// value
		hashedNodeValue := sha512.HashData(stateNode.Value)
		nodeHashInputData = append(nodeHashInputData, hashedNodeValue...)

	}

	// If the node has children, order the keys of child label names
	// For each child, recursively call this function to get back the hash of that child.
	// Append the child hashed key + hashed value to nodeHashInputData
	if len(stateNode.Children) > 0 {

		var childKeyHashToKeyString = make(map[string]string)
		var childKeyHashList []string

		// sort keys
		for keyString := range stateNode.Children {
			// hash the key (required to be there), track it, append to child key hash list
			keyHash := sha512.HashData([]byte(keyString))
			childKeyHashToKeyString[string(keyHash)] = keyString
			childKeyHashList = append(childKeyHashList, string(keyHash))
		}
		sort.Strings(childKeyHashList)

		// In sorted order, get each child and append to input data
		for _, childKeyHash := range childKeyHashList {

			// look up actual key name and get its child value (the node)
			actualKeyString := childKeyHashToKeyString[childKeyHash]
			childNodeBytes := stateNode.Children[actualKeyString]

			// decode child node
			var childNode protocol.CollaborationState_Node
			err := proto.Unmarshal(childNodeBytes, &childNode)
			if err != nil {
				return nil
			}

			// Recursively call this function on Node's child node
			// Append to input data
			childNodeStateHash := CalculateCollaborationStateHash(childNode)
			nodeHashInputData = append(nodeHashInputData, childNodeStateHash...)

		}

	}

	// Return the hash of this node's label_hash + value_hash + children_hashes
	return sha512.HashData(nodeHashInputData)

}
