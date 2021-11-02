package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customers struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestJsonObject(t *testing.T) {
	customer := Customers{
		FirstName:  "calvin",
		MiddleName: "wi",
		LastName:   "jaya",
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
