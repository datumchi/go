package address

import (
	"errors"
	"github.com/datumchi/go/generated/protocol"
	"regexp"
	"strings"
)

var addressRegexp = regexp.MustCompile(`\[(.*)\](.*)`)

func ToString(address protocol.Address) string {

	stringAddress := "[" + address.DescriptorReference
	if address.DescriptorPath == "" {
		stringAddress = stringAddress + "]"
	} else {
		stringAddress = stringAddress + address.DescriptorPath + "]"
	}

	stringAddress = stringAddress + address.Domain

	return stringAddress

}


func ToAddress(addressString string) (protocol.Address, error) {

	var address protocol.Address
	if(!addressRegexp.MatchString(addressString)) {
		return address, errors.New("Invalid address: " + addressString)
	}

	addressParts := addressRegexp.FindStringSubmatch(addressString)
	if len(addressParts) < 3 {
		return address, errors.New("Could not parse address: " + addressString)
	}

	address.Domain = addressParts[2]
	if strings.Contains(addressParts[1], "*") {

		descriptorParts := strings.Split(addressParts[1], "*")
		if len(descriptorParts) != 2 {
			return address, errors.New("Could not parse descriptor blob: " + addressParts[1])
		}

		address.DescriptorReference = descriptorParts[0]
		address.DescriptorPath = descriptorParts[1]

	} else {

		address.DescriptorReference = addressParts[1]

	}

	return address, nil


}
