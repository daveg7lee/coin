package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/daveg7lee/kangaroocoin/utils"
)

const (
	signature     string = "bb812d2805d02f78bf13565b5d4deada2997077417fbf62392ce1058ee22d9d7e128121763f48a95a9f3dc80d46e8e9fe1960f14f4118c32cf61c39a8d162237"
	privateKey    string = "30770201010420e53f9b70c2f7a00ae787969e4e86056efba651695ab49e9f9343b3c6598c1fd3a00a06082a8648ce3d030107a14403420004c7670a1ccb3d35d9124c76dc4d0af0851f1c6c2b96b4b7d6c0389c5410c57537f41621f888721a104e7d63c1d3fb7ba45083e8727cb2cdfe7bfb4fd222908934"
	hashedMessage string = "c47757abe4020b9168d0776f6c91617f9290e790ac2f6ce2bd6787c74ad88199"
)

func Start() {
	privateKeyAsBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	_, err = x509.ParseECPrivateKey(privateKeyAsBytes)
	utils.HandleErr(err)

	signatureAsBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := signatureAsBytes[:len(signatureAsBytes)/2]
	sBytes := signatureAsBytes[len(signatureAsBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}
