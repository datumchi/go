package domain

import (
	"github.com/datumchi/go/connect/dnsconfig"
	"github.com/datumchi/go/naming"
	"github.com/datumchi/go/utility/logger"
	domainverify "github.com/datumchi/go/verifier/domain"
)

func VerifyDomain(domain string) (naming.DomainDefinition, bool) {

	// get domain definition
	configRepo := dnsconfig.CreateNetworkAwareDNSConfigRepo()
	domainDefinition, err := configRepo.GetDomainDefinition(domain)
	if err != nil {
		logger.Fatalf("Fatal error while verifying domain %s:  %v", domain, err)
		return naming.DomainDefinition{}, false
	}

	if domainverify.VerifyDomainDefinition(domainDefinition) {
		return domainDefinition, true
	}

	return domainDefinition, false

}
