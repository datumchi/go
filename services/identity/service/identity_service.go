package service

import (
	"context"
	"github.com/datumchi/go/crypto/hsm"
	"github.com/datumchi/go/generated/protocol"
	"github.com/datumchi/go/services/identity/authlocalmemory"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/storage"
	"github.com/datumchi/go/utility/logger"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
)

func CreateIdentityService() (protocol.IdentityServicesServer, error) {

	config := configuration.CreateConfiguration()
	svc := IdentityService{
		IdentityDomain: config.Domain(),
		JWTAuthKey:     []byte(config.JWTKey()),
	}

	// Identity Authenticator
	switch strings.ToUpper(config.IdentityAuthenticationScheme()) {

	case configuration.IDENTITY_AUTHENTICATION_SCHEME_LOCAL_MEMORY:
		localMem := authlocalmemory.CreateLocalMemoryAuthenticator()
		svc.IdentityAuthScheme = localMem

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

	// return svc, nil
	return nil, nil

}

type IdentityAuthenticator interface {
	GetChallenge(authSubject string) string
	VerifyChallengeResponse(authSubject string, challengeResponse string) bool
}

type IdentityService struct {
	JWTAuthKey         []byte
	IdentityDomain     string
	IdentityStore      storage.BlobStore
	IdentityAuthScheme IdentityAuthenticator
}

func (identityService IdentityService) AuthenticateGetChallenge(ctx context.Context, authInfo *protocol.AuthenticationInfo) (*protocol.CommonResponse, error) {

	var response protocol.CommonResponse

	switch authInfo.Type {

	case protocol.AuthenticationInfo_OAUTH_JWT:
		authSubject := authInfo.Data // public key
		challenge := identityService.IdentityAuthScheme.GetChallenge(authSubject)
		response.IsOk = true
		response.ExtraInformation = challenge

	default:
		response.IsOk = false
		response.ExtraInformation = "Unsupported."

	}

	return &response, nil

}

func (identityService IdentityService) Authenticate(ctxm context.Context, authInfo *protocol.AuthenticationInfo) (*protocol.AuthenticationToken, error) {

	var response protocol.AuthenticationToken

	switch authInfo.Type {

	case protocol.AuthenticationInfo_OAUTH_JWT:
		var subjectInfo []string
		authData := authInfo.Data // public key ___!!!___ signature
		if strings.Contains(authData, "___!!!___") {
			subjectInfo = strings.Split(authData, "___!!!___")
		} else {
			return &response, nil
		}

		verifiedFlag := identityService.IdentityAuthScheme.VerifyChallengeResponse(subjectInfo[0], subjectInfo[1])
		if verifiedFlag {
			authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"public_key": subjectInfo[0],
			})

			authTokenSigned, err := authToken.SignedString(identityService.JWTAuthKey)
			if err != nil {
				return &response, nil
			}

			tokenItems := strings.Split(authTokenSigned, ".")
			response.Header = tokenItems[0]
			response.Payload = tokenItems[1]
			response.Signature = tokenItems[2]

		} else {
			return &response, nil
		}

	}

	return &response, nil

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

	return &protocol.CommonResponse{IsOk: true}, nil

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

	return &protocol.CommonResponse{IsOk: true}, nil

}

func (identityService IdentityService) DeliverMessage(context.Context, *protocol.IdentityMessage) (*protocol.CommonResponse, error) {
	return &protocol.CommonResponse{IsOk: true}, nil
}

func (identityService IdentityService) RetrieveMessages(context.Context, *protocol.AuthenticationToken) (*protocol.IdentityMessageBundle, error) {
	panic("implement me")
}
