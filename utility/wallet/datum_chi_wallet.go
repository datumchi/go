package main

import (
	"fmt"
	"github.com/datumchi/go/utility/logger"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"os"
)

func main() {

	mnemonic, err := hdwallet.NewMnemonic(256)
	if err != nil {
		logger.Errorf("Error generating mnemonic: %v", err)
		os.Exit(1)
	}

	// generate 10 ethereum accounts
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		logger.Errorf("Unable to create wallet using generated mnemonic: %v", err)
		os.Exit(1)
	}

	var accounts [10]string
	for i := 0 ; i < 10 ; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			logger.Errorf("Iteration %i resulted in a bad account derivation", err)
			os.Exit(1)
		}
		accounts[i] = account.Address.Hex()
	}


	logger.Infof("Mnemonic: %s", mnemonic)
	for i,a := range accounts {
		logger.Infof("\tAccount %v: %s", i, a)
	}




}