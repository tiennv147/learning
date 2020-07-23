package main

import (
	"context"
	"fmt"
)

type Verifier interface {
	Verify(context.Context)
}

type RequireBodyVerifier interface {
	RequireBody()
}

// RequireBody check if a verifier need body or not
func RequireBody(v interface{}) bool {
	needPhat, ok := v.(RequireBodyVerifier)
	fmt.Println(needPhat)
	if ok {
		return ok
	}
	return false
}


type onlyVerifier struct {
	good int
}

func (v *onlyVerifier) Verify(context.Context) {
}

type needBodyVerifier struct {
	best string
}

func (v *needBodyVerifier) Verify(context.Context) {
}

func (v *needBodyVerifier) RequireBody() {
}




func main() {
	var onlyKu Verifier
	var needKu Verifier

	onlyKu = &onlyVerifier{}
	needKu = &needBodyVerifier{}

	fmt.Printf("need 1: %v\n", RequireBody(onlyKu))
	fmt.Printf("need 2: %v\n", RequireBody(needKu))
}
