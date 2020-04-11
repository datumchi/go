package connect

import (
	"github.com/datumchi/go/generated/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type DomainNodeConnectFactory struct {

}

func (DomainNodeConnectFactory) NodeCollaborationClient(domain string) protocol.CollaborationServicesClient {

	// Get The Domain Configuration

	// Establish connection
	var opts []grpc.DialOption
	tlsCredentials := credentials.NewTLS(nil)
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.Dial(identityServiceUrl, opts...)
	if err != nil {
		return nil, errors.New("Client connection error:  " + err.Error())
	}

	client := protocol.NewIdentityServicesClient(conn)

	return client, nil
}

func (DomainNodeConnectFactory) NodeIdentityClient(domain string) protocol.IdentityServicesClient {
	panic("implement me")
}

