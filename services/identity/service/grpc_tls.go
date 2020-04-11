package service

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/utility/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"os"
)

func ConfigureGrpcTLSServer() (*grpc.Server, error) {

	// Configuration
	config := configuration.CreateConfiguration()
	var gs *grpc.Server

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(config.TlsServerCert(), config.TlsServerKey())
	if err != nil {
		logger.Fatalf("Could not load server key pair: %s", err)
		os.Exit(1)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(config.TlsCaCert())
	if err != nil {
		logger.Fatalf("Could not read ca certificate: %s", err)
		os.Exit(1)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		logger.Fatalf("Failed to append client certs")
		os.Exit(1)
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	gs = grpc.NewServer(grpc.Creds(creds))

	return gs, nil


}

