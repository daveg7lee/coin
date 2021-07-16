package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/daveg7lee/kangaroocoin/utils"
)

const (
	fileName string = "nomadcoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    string
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return
}

func addressFromKey(key *ecdsa.PrivateKey) string {

}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.address = addressFromKey(w.privateKey)
	}
	return w
}
