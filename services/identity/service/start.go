package service

import (
	"github.com/datumchi/go/naming"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/services/identity/domain"
	"github.com/datumchi/go/utility/logger"
	"os"
)

/** Startup Sequence:
 *    - Check to see if we should perform domain verification, including domain key and etc.
 *    - Check each blob storage configuration and configure.
 *    -
 */
func Start() {

	var domainDefinition naming.DomainDefinition
	var verifiedFlag bool = false
	config := configuration.CreateConfiguration()

	if config.VerifyDomain() {

		logger.Infof("(%s) Verifying Domain Configuration", config.Domain())
		domainDefinition, verifiedFlag = domain.VerifyDomain(config.Domain())
		if !verifiedFlag {
			logger.Fatalf("Unable to verify domain: %s", config.Domain())
			os.Exit(1)
		}

	}

	_ = domainDefinition

}
