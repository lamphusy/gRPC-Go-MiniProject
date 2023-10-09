package jwt

import (
	"fmt"
)

type Middleware struct{
	Validator
}

func NewMiddleware (publicKeyPath string) (*Middleware, error){
	validator, err := NewValidator(publicKeyPath)
	if err!=nil{
		return nil, fmt.Errorf("unable to create validator %w", err)
	}

	return &Middleware{
		Validator: *validator,
	},nil
}
