package dns

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

type DomainDefinition struct {
	Domain string
	PublicKey string
	Signature string
	IdentityServiceHost string
	IdentityServicePort uint16
	CollabServiceHost string
	CollabServicePort uint16
}

var publicKeyRegExp = regexp.MustCompile(`v=datumchi_pubkey; k=(.*)`)
var signatureRegExp = regexp.MustCompile(`v=datumchi_signature; s=(.*)`)


func GetDomainDefinition(domain string) (DomainDefinition, error) {

	domainDefinition := DomainDefinition{
		Domain:domain,
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
			err = parsePublicKey(entry, &domainDefinition)
			if err != nil {
				return domainDefinition, err
			}
		} else if strings.Contains(entry, "v=datumchi_signature") {
			err = parseSignature(entry, &domainDefinition)
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

func parsePublicKey(entry string, domainDefinition *DomainDefinition) error {

	submatches := publicKeyRegExp.FindStringSubmatch(entry)
	if len(submatches) != 2 {
		return errors.New("Unable to parse pubkey entry (TXT)")
	}

	domainDefinition.PublicKey = submatches[1]
	return nil

}

func parseSignature(entry string, domainDefinition *DomainDefinition) error {

	submatches := signatureRegExp.FindStringSubmatch(entry)
	if len(submatches) != 2 {
		return errors.New("Unable to parse signature entry (TXT)")
	}

	domainDefinition.Signature = submatches[1]
	return nil

}

