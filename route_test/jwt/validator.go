package jwt

import (
	"crypto"
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
)

type Validator struct {
	key crypto.PublicKey
}

func NewValidator(publicKeyPath string) (*Validator, error) {
	keyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read public key file: %w", err)
	}

	key, err := jwt.ParseEdPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse as ed private key: %w", err)
	}

	return &Validator{
		key: key,
	}, nil
}

func (v *Validator) GetToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return v.key, nil
		})
	if err != nil {
		return nil, fmt.Errorf("unable to parse token string: %w", err)
	}

	return token, nil
}
