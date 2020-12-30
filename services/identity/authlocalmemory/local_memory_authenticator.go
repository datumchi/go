package authlocalmemory

import (
	"github.com/datumchi/go/crypto/ed25519"
	"github.com/datumchi/go/encoding"
	"math/rand"
	"time"
)

type LocalMemoryAuthenticator struct {
	AuthenticationChallenges map[string]string
}

func CreateLocalMemoryAuthenticator() LocalMemoryAuthenticator {
	return LocalMemoryAuthenticator{}
}

func (localMemAuth LocalMemoryAuthenticator) GetChallenge(authSubject string) string {

	// generate random text
	challenge := randomString(32)
	localMemAuth.AuthenticationChallenges[authSubject] = challenge

	return challenge
}

func (localMemAuth LocalMemoryAuthenticator) VerifyChallengeResponse(authSubject string, challengeResponse string) bool {

	publicKey, err := encoding.Decode(authSubject)
	if err != nil {
		return false
	}

	challengeSignature, err := encoding.Decode(challengeResponse)
	if err != nil {
		return false
	}
	challengeText := localMemAuth.AuthenticationChallenges[authSubject]

	return ed25519.Verify(publicKey, challengeSignature, []byte(challengeText))

}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func randomStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(length int) string {
	return randomStringWithCharset(length, charset)
}
