package jwt

import (
	"crypto"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type Issuer struct {
	key crypto.PrivateKey
}

func NewIssuer(privateKeyPath string) (*Issuer, error) {
	keyByte, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(fmt.Errorf("unable to read private key file: %w", err))
	}
	key, err := jwt.ParseEdPrivateKeyFromPEM(keyByte)
	if err != nil {
		return nil, fmt.Errorf("unable to parse as ed private key: %w", err)
	}
	return &Issuer{
		key: key,
	}, nil
}

func (i *Issuer) IssuerToken(user string, role []string) (string, error) {
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, jwt.MapClaims{
		//
	})
	tokenString, err := token.SignedString(i.key)
	if err != nil {
		return "", fmt.Errorf("unable to sign token: %w", err)
	}
	return tokenString, nil
}
