package dnsconfig

import "github.com/datumchi/go/naming"

type DNSConfigurationRepository interface {

	GetDomainDefinition(domain string) (naming.DomainDefinition, error)

}
