package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "goTemp/user/proto"
	"time"
)

//key: Encoding key
var key []byte

//TokenValidityPeriod is the length of time a token can be used
const TokenValidityPeriod = time.Hour * 24

//ClaimIssuer ia a string representing the claim issuer
const ClaimIssuer = "goTemp.usersrv"

//getKeyFromVault  checks to see if the key is populated and returns it.If key is empty, it is fetched from an external source
func getKeyFromVault() ([]byte, error) {
	if key == nil {
		//TODO: read key from Vault
		key = []byte("LOOKMEUPINEXTERNALSYSTEM")
	}
	return key, nil
}

//MyCustomClaims ia the structure of the claim to be used for the token
type MyCustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

//TokenService holds all methods related to tokens
type TokenService struct{}

//Decode a token and check its validity
func (ts *TokenService) Decode(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			key, err := getKeyFromVault()
			if err != nil {
				return nil, err
			}
			return key, nil
		})

	if token == nil {
		return nil, jwt.NewValidationError(glErr.AuthNilToken(), jwt.ValidationErrorUnverifiable)
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//Encode creates a new token based on a custom claim and return the signed string
func (ts *TokenService) Encode(user *pb.User) (string, error) {

	// Build Claim
	currentTime := time.Now()
	expireTime := currentTime.Add(TokenValidityPeriod)

	claims := MyCustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  currentTime.Unix(),
			Issuer:    ClaimIssuer,
		},
	}

	//Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Build signed string with our key
	key, err := getKeyFromVault()
	if err != nil {
		return "", err
	}

	ss, err := token.SignedString(key)

	return ss, nil
}
