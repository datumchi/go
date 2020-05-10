package cmd

import (
	"github.com/datumchi/go/crypto/ed25519"
	"github.com/datumchi/go/crypto/sha512"
	"github.com/datumchi/go/encoding"
	"github.com/datumchi/go/utility/logger"
	"github.com/spf13/cobra"
)

var verifydomainsigCmd = &cobra.Command{
	Use:   "verifydomainsig domain pubkey signature",
	Short: "verifydomainsig",
	Long:  `Verify a Signature for a Domain`,
	Run:   VerifyDomainSignature,
	Args: cobra.ExactArgs(3),
}



func init() {

	rootCmd.AddCommand(verifydomainsigCmd)
}


func VerifyDomainSignature(cmd *cobra.Command, args []string) {

	domain := args[0]
	pubkeyEncoded := args[1]
	signatureEncoded := args[2]

	// hash the domain
	domainHash := sha512.HashData([]byte(domain))
	pubkey, _ := encoding.Decode(pubkeyEncoded)
	signature, _ := encoding.Decode(signatureEncoded)

	// verify
	if ed25519.Verify(pubkey, signature, domainHash) {
		logger.Infof("%s - VERIFIED", domain)
	} else {
		logger.Infof("%s - FAILED", domain)
	}



}
