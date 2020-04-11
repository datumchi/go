package dns

import (
	"errors"
	"github.com/datumchi/go/naming"
	"net"
	"regexp"
	"strings"
)

type DNSDomainDefinition struct {
	PublicKeyRegExp *regexp.Regexp
	SignatureRegExp *regexp.Regexp
}

func CreateDNSDomainDefinition() DNSDomainDefinition {

	domainDefinition := DNSDomainDefinition{
		PublicKeyRegExp: regexp.MustCompile(`v=datumchi_pubkey; k=(.*)`),
		SignatureRegExp: regexp.MustCompile(`v=datumchi_signature; s=(.*)`),
	}

	return domainDefinition

}

func (dd DNSDomainDefinition) GetDomainDefinition(domain string) (naming.DomainDefinition, error) {

	domainDefinition := naming.DomainDefinition{
		Domain: domain,
	}

	textEntries, err := net.LookupTXT(domain)
	if err != nil {
		return domainDefinition, err
	}

	_, srvIdentityEntries, err := net.LookupSRV("datumchiidentity", "tcp", domain)
	if err != nil {
		return domainDefinition, err
	}

	_, srvCollabEntries, err := net.LookupSRV("datumchicollaboration", "tcp", domain)
	if err != nil {
		return domainDefinition, err
	}

	for _, entry := range textEntries {

		if strings.Contains(entry, "v=datumchi_pubkey") {
			err = dd.parsePublicKey(entry, &domainDefinition)
			if err != nil {
				return domainDefinition, err
			}
		} else if strings.Contains(entry, "v=datumchi_signature") {
			err = dd.parseSignature(entry, &domainDefinition)
			if err != nil {
				return domainDefinition, err
			}
		} else {
			return domainDefinition, errors.New("Proper Text Entries Not Found")
		}

	}

	if len(srvIdentityEntries) > 0 {
		domainDefinition.IdentityServiceHost = srvIdentityEntries[0].Target
		domainDefinition.IdentityServicePort = srvIdentityEntries[0].Port
	} else {
		return domainDefinition, errors.New("No SRV identity entries found")
	}

	if len(srvCollabEntries) > 0 {
		domainDefinition.CollabServiceHost = srvCollabEntries[0].Target
		domainDefinition.CollabServicePort = srvCollabEntries[0].Port
	} else {
		return domainDefinition, errors.New("No SRV collaboration entries found")
	}

	return domainDefinition, nil

}

func (dd DNSDomainDefinition) parsePublicKey(entry string, domainDefinition *naming.DomainDefinition) error {

	submatches := dd.PublicKeyRegExp.FindStringSubmatch(entry)
	if len(submatches) != 2 {
		return errors.New("Unable to parse pubkey entry (TXT)")
	}

	domainDefinition.PublicKey = submatches[1]
	return nil

}

func (dd DNSDomainDefinition) parseSignature(entry string, domainDefinition *naming.DomainDefinition) error {

	submatches := dd.SignatureRegExp.FindStringSubmatch(entry)
	if len(submatches) != 2 {
		return errors.New("Unable to parse signature entry (TXT)")
	}

	domainDefinition.Signature = submatches[1]
	return nil

}
