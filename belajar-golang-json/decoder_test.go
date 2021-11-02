package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	reader, _ := os.Open("customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customers{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("example.json")
	encoder := json.NewEncoder(writer)

	customer := Customers{
		FirstName:  "calvin",
		MiddleName: "wi",
		LastName:   "jaya",
	}

	encoder.Encode(customer)
	fmt.Println(customer)
}
