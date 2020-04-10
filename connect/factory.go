package connect

import "github.com/datumchi/go/generated/protocol"

type NodeConnectFactory interface {

	NodeCollaborationClient(domain string) protocol.CollaborationServicesClient
	NodeIdentityClient(domain string) protocol.IdentityServicesClient

}
