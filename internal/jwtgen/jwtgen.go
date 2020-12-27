package jwtgen

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Username       string `json:"username"`
	OwnedFarmIDs   []int  `json:"ownedFarmIds"`
	WorkingFarmIDs []int  `json:"workingFarmIds"`
	jwt.StandardClaims
}

// IJWTGenerator is an interface of JWT Generator
type IJWTGenerator interface {
	Gen(username string, ownedFarmIDs []int, workingFarmIDs []int) (jwt string, err error)
}

// JWTGenerator is an implementation of JWT Generator interface
type JWTGenerator struct {
	secret string
}

// Gen is a method for generating jwt token
func (g JWTGenerator) Gen(username string, ownedFarmIDs []int, workingFarmIDs []int) (t string, err error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		ownedFarmIDs,
		workingFarmIDs,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte(g.secret))
	if err != nil {
		return t, err
	}
	return t, nil
}

// NewJWTGenerator is a factory method of jwt generator
func NewJWTGenerator(secret string) JWTGenerator {
	return JWTGenerator{
		secret: secret,
	}
}
