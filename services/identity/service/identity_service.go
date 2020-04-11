package service

import (
	"context"
	"github.com/datumchi/go/crypto/hsm"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/storage"
	"github.com/datumchi/go/utility/logger"
	"os"
	"strings"
)

func CreateIdentityService() (protocol.IdentityServicesServer, error) {

	config := configuration.CreateConfiguration()
	svc := IdentityService{
		IdentityDomain:config.Domain(),
		JWTAuthKey:[]byte(config.JWTKey()),
		AuthenticationChallenges:make(map[string]string),
	}


	// Device Storage Options
	switch strings.ToUpper(config.DeviceStore()) {

	case configuration.STORE_TYPE_FILE:
		fileBlobStore, err := storage.CreateFileBlobStore(config.DeviceStoreFileBaseDir())
		if err != nil {
			logger.Fatalf("Fatal error while initializing a file blob store:  %v", err)
			os.Exit(1)
		}
		svc.DeviceStore = fileBlobStore


	default:
		logger.Fatalf("Option for device store (%v) unavailable", config.DeviceStore())
		os.Exit(1)


	}


	// Identity Storage Options
	switch strings.ToUpper(config.IdentityStore()) {

	case configuration.STORE_TYPE_FILE:
		var fileBlobStore storage.BlobStore
		var err error
		if strings.ToUpper(config.IdentityStoreFileSecurityModule()) == "SIMPLE" {
			key := hsm.NewEncryptionKey()
			securityModule, err := hsm.CreateSimpleSecurityModule(key)
			if err != nil {
				logger.Fatalf("Unable to create a simple tls module: %v", err)
				os.Exit(1)
			}
			fileBlobStore, err = storage.CreateSecureFileBlobStore(config.IdentityStoreFileBaseDir(), &securityModule)
		} else {
			fileBlobStore, err = storage.CreateFileBlobStore(config.IdentityStoreFileBaseDir())
		}
		if err != nil {
			logger.Fatalf("Fatal error while initializing a file blob store: %v", err)
			os.Exit(1)
		}
		svc.IdentityStore = fileBlobStore

	default:
		logger.Fatalf("Option for identity store (%v) unavailable", config.IdentityStore())
		os.Exit(1)



	}

	return svc, nil

}



type IdentityService struct {

	AuthenticationChallenges map[string]string

	JWTAuthKey []byte
	IdentityDomain string
	IdentityStore storage.BlobStore
	DeviceStore storage.BlobStore

}

func (identityService IdentityService) RegisterDevice(ctx context.Context, deviceInfo *protocol.DeviceInfo) (*protocol.CommonResponse, error) {

	//if !device.ValidateDevice(identityService.IdentityDomain, deviceInfo) {
		//	return &protocol.CommonResponse{
		//		IsOk:                 false,
		//		ExtraInformation:     "Unable to validate device info given",
		//	}, nil
		//}
		//
		//deviceInfoWrapper := device.WrapDeviceInfo("", deviceInfo)
		//if !device.SaveDevice(identityService.DeviceStore, deviceInfoWrapper) {
		//	return &protocol.CommonResponse{
		//		IsOk:                 false,
		//		ExtraInformation:     "Unable to save device info",
		//	}, nil
		//}

	return &protocol.CommonResponse{
		IsOk:                 true,
	}, nil

}

func (identityService IdentityService) AuthenticateGetChallenge(ctx context.Context, deviceInfo *protocol.DeviceInfo) (*protocol.CommonResponse, error) {

	var response protocol.CommonResponse
	//if identityService.AuthenticationChallenges[deviceInfo.DevicePublicKey] != "" {
	//	response = protocol.CommonResponse{
	//		IsOk:                 true,
	//		ExtraInformation:     identityService.AuthenticationChallenges[deviceInfo.DevicePublicKey],
	//	}
	//
	//	return &response, nil
	//}
	//
	//challenge := authentication.GenerateChallenge()
	//identityService.AuthenticationChallenges[deviceInfo.DevicePublicKey] = challenge
	//
	//response = protocol.CommonResponse{
	//	IsOk:             true,
	//	ExtraInformation: challenge,
	//}

	return &response, nil


}

func (identityService IdentityService) AuthenticateDevice(ctxm context.Context, deviceInfo *protocol.DeviceInfo) (*protocol.AuthenticationToken, error) {

	// does this device exist already?
	//deviceData, err := identityService.DeviceStore.Get(deviceInfo.DevicePublicKey)
	//if err != nil || deviceData == nil {
	//	return nil, errors.New("Unable to authenticate device")
	//}
	//
	//challenge := identityService.AuthenticationChallenges[deviceInfo.DevicePublicKey]
	//delete(identityService.AuthenticationChallenges, deviceInfo.DevicePublicKey)
	//return authentication.AuthenticateDevice(deviceInfo, identityService.IdentityDomain, challenge, identityService.JWTAuthKey)
	return nil, nil

}

func (identityService IdentityService) EstablishIdentity(ctx context.Context, request *protocol.EstablishIdentityAuthenticatedRequest) (*protocol.CommonResponse, error) {

	//authenticatedFlag, devicePublicKey := authentication.VerifyAuthenticationToken(*request.AuthToken, identityService.JWTAuthKey)
	//if !authenticatedFlag {
	//	return &protocol.CommonResponse{IsOk:false, ExtraInformation:"Unable to authenticate device to establish identity"}, nil
	//}
	//
	//// Verify correctly attested attributes
	//verifyResult := iv.VerifyIdentityAttributes(*request.Identity)
	//if len(verifyResult.InvalidAttestedAttributes) > 0 {
	//	var invalidAttestations string
	//	for k,_ := range verifyResult.InvalidAttestedAttributes {
	//		invalidAttestations = invalidAttestations + " " + k
	//	}
	//	return &protocol.CommonResponse{IsOk:false, ExtraInformation:"Invalid attribute attestation: " + invalidAttestations}, nil
	//}
	//
	//
	//var associatedDevices = []string {devicePublicKey}
	//idWrapper := identity.WrapIdentity(*request.Identity,associatedDevices)
	//
	//savedOk := identity.SaveIdentity(identityService.IdentityStore, idWrapper)
	//if !savedOk {
	//	return &protocol.CommonResponse{IsOk:false, ExtraInformation:"Unable to save identity."}, nil
	//}

	return &protocol.CommonResponse{IsOk:true}, nil

}

func (identityService IdentityService) GetIdentity(ctx context.Context, identityAddress *protocol.Address) (*protocol.Identity, error) {

	var ident protocol.Identity

	//base32identity := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString([]byte(address.ToString(*identityAddress)))
	//idWrapper, err := identity.GetIdentity(identityService.IdentityStore, base32identity)
	//if err != nil {
	//	return &protocol.Identity{}, nil
	//}
	//
	//err = proto.Unmarshal(idWrapper.IdentityRaw, &ident)
	//if err != nil {
	//	return &protocol.Identity{}, nil
	//}

	return &ident, nil

}

func (identityService IdentityService) AttestIdentityAttribute(ctx context.Context, request *protocol.AttestIdentityAttributeAuthenticatedRequest) (*protocol.CommonResponse, error) {

	//authenticatedFlag, _ := authentication.VerifyAuthenticationToken(*request.AuthToken, identityService.JWTAuthKey)
	//if !authenticatedFlag {
	//	return &protocol.CommonResponse{IsOk:false, ExtraInformation:"Unable to authenticate device to establish identity"}, nil
	//}
	//
	//attestOk, err := identity.AttestIdentityAttribute(identityService.IdentityStore, *request.IdentityAddress, request.IdentityAttestation)
	//if err != nil || !attestOk {
	//	return &protocol.CommonResponse{IsOk:false, ExtraInformation:"Attestation failed"}, nil
	//}

	return &protocol.CommonResponse{IsOk:true}, nil

}

func (identityService IdentityService) DeliverMessage(context.Context, *protocol.IdentityMessage) (*protocol.CommonResponse, error) {
	return &protocol.CommonResponse{IsOk:true}, nil
}

func (identityService IdentityService) RetrieveMessages(context.Context, *protocol.AuthenticationToken) (*protocol.IdentityMessageBundle, error) {
	panic("implement me")
}



