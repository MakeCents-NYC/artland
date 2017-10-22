package oauth

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

type OAuth struct {
	ConsumerKey  string
	PrivateKey   interface{}
	Pem          *pem.Block
	Cert         *x509.Certificate
	KeyAlias     string
	KeyStorePass string
}

func NewOAuth(_cKey, p12, kA, kSP string) (*OAuth, error) {

	result := &OAuth{
		ConsumerKey:  _cKey,
		KeyAlias:     kA,
		KeyStorePass: kSP,
	}

	p12data, err := ioutil.ReadFile(p12)
	if err != nil {
		return nil, err
	}

	result.PrivateKey, result.Cert, err = pkcs12.Decode(p12data, kSP)
	if err != nil {
		return nil, err
	}

	block, err := pkcs12.ToPEM(p12data, kSP)
	if err != nil {
		return nil, err
	}

	return result, nil
}
