package main

import (
	_ "embed"
	"errors"
	"net/http"

	"grpc-test/route_test/jwt"
)



type AuthService struct {
	issuer *jwt.Issuer
}

func NewAuthService(issuer *jwt.Issuer) (*AuthService, error) {
	if issuer == nil {
		return nil, errors.New("issuer is required")
	}

	return &AuthService{
		issuer: issuer,
	}, nil
}

func (a *AuthService) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// user, pass, ok := r.BasicAuth()
	// if !ok {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("missing basic auth"))
	// 	return
	// }

	// if user != "admin" || pass != "pass" {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte("invalid credentials"))
	// 	return
	// }

	// tokenString, err := a.issuer.IssuerToken("admin", []string{"admin", "basic"})
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("unable to issue token:" + err.Error()))
	// 	return
	// }

	// _, _ = w.Write([]byte(tokenString + "\n"))
}
