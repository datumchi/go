package service

import (
	"fmt"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/naming"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/services/identity/domain"
	"github.com/datumchi/go/utility/logger"
	"net"
	"os"
)

/** Startup Sequence:
 *    - Check to see if we should perform domain verification, including domain key and etc.
 *    - Check each blob storage configuration and configure identity service
 *    - Configure gRPC over TLS
 */
func Start() {

	// Check to see if we should perform domain verification, including domain key and etc.
	var domainDefinition naming.DomainDefinition
	var verifiedFlag bool = false
	config := configuration.CreateConfiguration()

	if config.VerifyDomain() {

		logger.Infof("(%s) Verifying Domain Configuration", config.Domain())
		domainDefinition, verifiedFlag = domain.VerifyDomain(config.Domain())
		if !verifiedFlag {
			logger.Fatalf("(%s) Unable to verify domain :(", config.Domain())
			os.Exit(1)
		}

		logger.Infof("(%s) Domain Verified", config.Domain())

	} else {

		// populate domain definition here


	}
	logger.Infof("(%s) Domain Public Key:  %s", config.Domain(), domainDefinition.PublicKey)
	logger.Infof("(%s)  Domain Signature:  %s", config.Domain(), domainDefinition.Signature)


	// Check each blob storage configuration and configure identity service
	identityServiceServer, err := CreateIdentityService()
	if err != nil {
		logger.Fatalf("(%s) Unable to create identity service: %v", config.Domain(), err)
		os.Exit(1)
	}

	// Configure gRPC over TLS
	grpcServer, err := ConfigureGrpcTLSServer()
	if err != nil {
		logger.Fatalf("(%s) Unable to create gRPC server: %v", config.Domain(), err)
		os.Exit(1)
	}
	protocol.RegisterIdentityServicesServer(grpcServer, identityServiceServer)

	logger.Infof("Starting Identity Service")
	listenSocket, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.ServiceHost(), config.ServicePort()))
	if err != nil {
		logger.Fatalf("(%s) Unable to listen to socket: %v", config.Domain(), err)
		os.Exit(1)
	}

	grpcServer.Serve(listenSocket)

}
