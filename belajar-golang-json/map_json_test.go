package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"apple macbook pro","price":"20000000"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestJsonMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "mbp m1",
		"price": 2000000,
	}

	bytes, _ := json.Marshal(product)

	fmt.Println(string(bytes))
}
