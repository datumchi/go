package service

import (
	context "context"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/services/identity/configuration"
)

func CreateHeartbeatOrdererService() (protocol.OrderingServicesServer, error) {

	config := configuration.CreateConfiguration()
	svc := HeartbeatOrderingService {

	}

	//
	//
	//// Device Storage Options
	//switch strings.ToUpper(config.DeviceStore()) {
	//
	//case configuration.STORE_TYPE_FILE:
	//	fileBlobStore, err := storage.CreateFileBlobStore(config.DeviceStoreFileBaseDir())
	//	if err != nil {
	//		logger.Fatalf("Fatal error while initializing a file blob store:  %v", err)
	//		os.Exit(1)
	//	}
	//	svc.DeviceStore = fileBlobStore
	//
	//
	//default:
	//	logger.Fatalf("Option for device store (%v) unavailable", config.DeviceStore())
	//	os.Exit(1)
	//
	//
	//}
	//
	//
	//// Identity Storage Options
	//switch strings.ToUpper(config.IdentityStore()) {
	//
	//case configuration.STORE_TYPE_FILE:
	//	var fileBlobStore storage.BlobStore
	//	var err error
	//	if strings.ToUpper(config.IdentityStoreFileSecurityModule()) == "SIMPLE" {
	//		key := hsm.NewEncryptionKey()
	//		securityModule, err := hsm.CreateSimpleSecurityModule(key)
	//		if err != nil {
	//			logger.Fatalf("Unable to create a simple tls module: %v", err)
	//			os.Exit(1)
	//		}
	//		fileBlobStore, err = storage.CreateSecureFileBlobStore(config.IdentityStoreFileBaseDir(), &securityModule)
	//	} else {
	//		fileBlobStore, err = storage.CreateFileBlobStore(config.IdentityStoreFileBaseDir())
	//	}
	//	if err != nil {
	//		logger.Fatalf("Fatal error while initializing a file blob store: %v", err)
	//		os.Exit(1)
	//	}
	//	svc.IdentityStore = fileBlobStore
	//
	//default:
	//	logger.Fatalf("Option for identity store (%v) unavailable", config.IdentityStore())
	//	os.Exit(1)
	//
	//
	//
	//}
	//
	//return svc, nil

	return svc, nil
}



type HeartbeatOrderingService struct {



}

func (hbos HeartbeatOrderingService) CreateCollaboration(ctx context.Context, collaboration *protocol.Collaboration) (*protocol.CommonResponse, error) {
	panic("implement me")
}

func (hbos HeartbeatOrderingService) DeliverCollaboration(ctx context.Context, collaboration *protocol.Collaboration) (*protocol.CommonResponse, error) {
	panic("implement me")
}

func (hbos HeartbeatOrderingService) OrderChangeInstruction(ctx context.Context, instruction *protocol.CollaborationChangeInstruction) (*protocol.CommonResponse, error) {
	panic("implement me")
}
