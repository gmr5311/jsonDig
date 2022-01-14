package jsondig

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name    string  `json:"name,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	Number int    `json:"number,omitempty"`
	Street string `json:"street,omitempty"`
}

func TestJsonDig(t *testing.T) {
	test := Test{
		Name: "Bob",
		Address: Address{
			Number: 1234,
			Street: "Main",
		},
	}
	thing, err := json.Marshal(test)

	assert.Equal(t, true, FindInJSON(string(thing), "name", "Bob"))
	assert.Equal(t, nil, err)
}
