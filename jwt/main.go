package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// LoadRSAPrivateKeyFromDisk ...
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

// VerifyToken ...
func VerifyToken(tokenStr string, key interface{}) bool {
	// Step 1: parse token by public rsa key, which is load by provider setting in the endpoint
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return token.Valid
}

// LoadRSAPublicKeyFromDisk ...
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

// MakeSampleToken ...
func MakeSampleToken(c jwt.Claims, key interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, e := token.SignedString(key)

	if e != nil {
		panic(e.Error())
	}

	return s
}

// LoadRSAPublicKeyFromString ...
func LoadRSAPublicKeyFromString(keyData string) *rsa.PublicKey {
	key, e := jwt.ParseRSAPublicKeyFromPEM([]byte(keyData))
	if e != nil {
		panic(e.Error())
	}
	return key
}

func main() {
	fmt.Println("-----")
	// privateKey := LoadRSAPrivateKeyFromDisk("/Users/tien.nguyenvan/Documents/ssh-key-grpc-proxy/grpc-proxy.key")
	// token := MakeSampleToken(jwt.MapClaims{
	// 	"authorized": "true",
	// 	"iss":        "careem",
	// 	"sub":        "careem",
	// 	"aud":        "PASSENGER",
	// 	"exp":        time.Now().Add(time.Hour * 24).Unix(),
	// }, privateKey)
	// fmt.Println(token)
	// fmt.Println("-----")

	privateKey2 := LoadRSAPrivateKeyFromDisk("/Users/tien.nguyenvan/Documents/ssh-key-test/test-uber.key")
	token2 := MakeSampleToken(jwt.StandardClaims{
		Audience:  "ENTERPRISE",
		ExpiresAt: time.Now().Add(time.Hour * 240).Unix(),
		Id:        "d34192bb-dd00-4de6-b333-4149e3c183a0",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "Uber",
		NotBefore: time.Now().Unix(),
		Subject:   "d34192bb-dd00-4de6-b333-4149e3c183a0",
	}, privateKey2)
	fmt.Println(token2)

	fmt.Println(VerifyToken(token2, privateKey2.Public()))
}
