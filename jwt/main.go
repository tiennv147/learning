package main

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func MakeSampleToken(c jwt.Claims, key interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}


func main() {
	privateKey := LoadRSAPrivateKeyFromDisk("/Users/tien.nguyenvan/Documents/ssh-key-grpc-proxy/grpc-proxy.key")
	token := MakeSampleToken(jwt.MapClaims{
		"authorized": "true",
		"iss": "carem",
		"sub": "carem",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}, privateKey)
	fmt.Println(token)
}