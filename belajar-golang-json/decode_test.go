package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	FirstName, MiddleName, LastName string
}

func TestDecode(t *testing.T) {
	jsonRequest := `{"FirstName":"calvin","MiddleName":"wi","LastName":"jaya"}`
	jsonBytes := []byte(jsonRequest)

	customers := &Person{}
	err := json.Unmarshal(jsonBytes, customers)
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
