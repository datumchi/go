package naming



type DomainDefinition struct {
	Domain string
	PublicKey string
	Signature string
	IdentityServiceHost string
	IdentityServicePort uint16
	CollabServiceHost string
	CollabServicePort uint16
}


type NamingServiceConfiguration interface {

	GetDomainDefinition(domain string) (DomainDefinition, error)

}