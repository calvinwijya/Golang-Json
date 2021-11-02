package belajargolangjson

import (
	"encoding/json"
	"testing"
)

func LogJson(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	println(string(byte))
}

func TestEncode(t *testing.T) {
	LogJson("calvin")
	LogJson(31)
	LogJson(true)
	LogJson([]string{"calvin", "wijaya"})
}
