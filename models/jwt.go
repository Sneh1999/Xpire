package models

import "github.com/dgrijalva/jwt-go"

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	ID string
	jwt.StandardClaims
}
