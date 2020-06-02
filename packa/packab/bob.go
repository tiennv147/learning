package packtuktuk

import "fmt"

type Bob struct {
	Test int
	Tit  string
}

func (b Bob) Name() string {
	return "Bob"
}

func HelloBob() {
	fmt.Println("HelloBob")
}
