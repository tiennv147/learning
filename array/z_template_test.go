package array

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Dog struct {
	ID    int
	Name  string
	Breed string
}

func (d Dog) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func TestTemplate1(t *testing.T) {
	//input := []int{1, 3, -1}
	//expected := 5
	//result := template(input)
	//

	dog := Dog{1, "bowser", "husky"}

	bytes, err := json.Marshal(dog)

	fmt.Println(string(bytes))

	assert.Nil(t, err)

}
