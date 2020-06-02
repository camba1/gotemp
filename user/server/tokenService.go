package main

import (
	"github.com/dgrijalva/jwt-go"
	pb "goTemp/user/proto"
	"time"
)

var key []byte

const TokenValidityPeriod = time.Hour * 24

//getKeyFromVault: Checks to see if the key is populated and returns it.If key is empty, it is fetched from an external source
func getKeyFromVault() ([]byte, error) {
	if key == nil {
		key = []byte("LOOKMEUPINEXTERNALSYSTEM")
	}
	return key, nil
}

type MyCustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct{}

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
		return nil, jwt.NewValidationError("Invalid nil user token", jwt.ValidationErrorUnverifiable)
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (ts *TokenService) Encode(user *pb.User) (string, error) {

	// Build Claim
	currentTime := time.Now()
	expireTime := currentTime.Add(TokenValidityPeriod)

	claims := MyCustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  currentTime.Unix(),
			Issuer:    "goTemp.usersrv",
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
