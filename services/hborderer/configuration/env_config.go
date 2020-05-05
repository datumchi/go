package configuration

import (
	"crypto/rand"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Loaded bool
}

var (
	SERVICE_HOST = "SERVICE_HOST"
	SERVICE_PORT = "SERVICE_PORT"
	TLS_SERVER_CERT = "TLS_SERVER_CERT"
	TLS_SERVER_KEY = "TLS_SERVER_KEY"
	TLS_CA_CERT = "TLS_CA_CERT"

	DOMAIN = "DOMAIN"
	VERIFY_DOMAIN="VERIFY_DOMAIN"

	JWT_KEY = "JWT_KEY"

	loadedConfiguration Configuration
)


func CreateConfiguration() Configuration {

	if loadedConfiguration.Loaded {
		return loadedConfiguration
	}

	loadedConfiguration = Configuration{}

	viper.SetEnvPrefix("DATUMCHI")
	viper.AutomaticEnv()

	viper.SetDefault(SERVICE_HOST, "localhost")
	viper.SetDefault(SERVICE_PORT, "17117")
	viper.SetDefault(DOMAIN, "localhost")
	viper.SetDefault(VERIFY_DOMAIN, "true")

	randomJwtKey := make([]byte, 20)
	rand.Read(randomJwtKey)
	viper.SetDefault(JWT_KEY, string(randomJwtKey))

	loadedConfiguration.Loaded = true
	return loadedConfiguration

}


func (c Configuration) ServiceHost() string {
	return viper.GetString(SERVICE_HOST)
}

func (c Configuration) ServicePort() string {
	return viper.GetString(SERVICE_PORT)
}

func (c Configuration) TlsServerCert() string {
	return viper.GetString(TLS_SERVER_CERT)
}

func (c Configuration) TlsServerKey() string {
	return viper.GetString(TLS_SERVER_KEY)
}

func (c Configuration) TlsCaCert() string {
	return viper.GetString(TLS_CA_CERT)
}

func (c Configuration) Domain() string {
	return viper.GetString(DOMAIN)
}

func (c Configuration) JWTKey() string {
	return viper.GetString(JWT_KEY)
}

func (c Configuration) VerifyDomain() bool {

	if strings.ToUpper(viper.GetString(VERIFY_DOMAIN)) == "TRUE" {
		return true
	}

	return false

}


