package main

import (
	"fmt"
)

type EmptyStruct struct{}

var (
	p1e = &EmptyStruct{}
	p2e = &EmptyStruct{}
)

type NotEmptyStruct struct{ val int }

var (
	p1ne = &NotEmptyStruct{}
	p2ne = &NotEmptyStruct{}
)

func main() {
	//https://golang.org/ref/spec#Size_and_alignment_guarantees
	//go run -gcflags '-m' compare.go

	// Two distinct zero-size variables may have the same address in memory.
	fmt.Println("[EMPTY][GLOBAL][same address]", p1e == p2e)
	// Pointers to distinct zero-size variables may or may not be equal.
	fmt.Println("[EMPTY][INLINE][may or may not be equal]", &EmptyStruct{} == &EmptyStruct{})

	fmt.Println("[NOT-EMPTY][INLINE][address not equal]", &NotEmptyStruct{} == &NotEmptyStruct{})
	fmt.Println("[NOT-EMPTY][GLOBAL][address not equal]", p1ne == p2ne)
}
