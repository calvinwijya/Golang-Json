package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country string
	PostalCode      int
}

type Player struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

func TestJsonArrayEncode(t *testing.T) {
	customer := Player{
		FirstName:  "calvin",
		MiddleName: "wi",
		LastName:   "jaya",
		Hobbies:    []string{"game", "judi", "coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"calvin","MiddleName":"wi","LastName":"jaya","Hobbies":["game","judi","coding"]}`
	jsonBytes := []byte(jsonString)

	Player1 := &Player{}

	err := json.Unmarshal(jsonBytes, Player1)
	if err != nil {
		panic(err)
	}
	fmt.Println(Player1)
}

func TestJsonArrayComplexEncode(t *testing.T) {
	customer := Player{
		FirstName: "calvin",
		Addresses: []Address{
			{
				Street:     "baloi",
				Country:    "indonesia",
				PostalCode: 29444,
			},
			{
				Street:     "permata",
				Country:    "indonesia",
				PostalCode: 29441,
			},
		},
	}

	bytes, _ := json.Marshal(customer)

	fmt.Print(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"calvin","MiddleName":"","LastName":"","Hobbies":null,"Addresses":[{"Street":"baloi","Country":"indonesia","PostalCode":29444},{"Street":"permata","Country":"indonesia","PostalCode":29441}]}`
	jsonBytes := []byte(jsonString)

	Player1 := &Player{}

	err := json.Unmarshal(jsonBytes, Player1)
	if err != nil {
		panic(err)
	}
	fmt.Println(Player1)
}

func TestJsonArray(t *testing.T) {
	jsonString := `[{"Street":"baloi","Country":"indonesia","PostalCode":29444},{"Street":"permata","Country":"indonesia","PostalCode":29441}]}`
	jsonBytes := []byte(jsonString)

	Addresses := &[]Address{}

	err := json.Unmarshal(jsonBytes, Addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(Addresses)
}
