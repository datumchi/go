package cmd

import (
	"crypto/rand"
	"github.com/datumchi/go/crypto/ed25519"
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/encoding"
	"github.com/datumchi/go/utility/logger"
	"github.com/spf13/cobra"
)

var keygenCmd = &cobra.Command{
	Use:   "domainkeygen",
	Short: "domainkeygen",
	Long:  `Generate a pair of EdDSA (25519 Curve) keys`,
	Run:   GenKeys,
	Args: cobra.ExactArgs(1),
}



func init() {

	rootCmd.AddCommand(keygenCmd)
}


func GenKeys(cmd *cobra.Command, args []string) {

	domainName := args[0]
	hashedDomainName := sha512.HashData([]byte(domainName))

	publicKey, privateKey := ed25519.GenerateKeypair(rand.Reader)
	encodedPublicKey := encoding.Encode(publicKey)
	encodedPrivateKey := encoding.Encode(privateKey)

	signature := ed25519.Sign(privateKey, hashedDomainName)
	signatureEncoded := encoding.Encode(signature)

	logger.Infof("Public Key:   %s", encodedPublicKey)
	logger.Infof("Private Key:  %s", encodedPrivateKey)
	logger.Infof("Signature:    %s", signatureEncoded)

}