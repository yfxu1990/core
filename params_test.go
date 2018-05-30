package main

import (
	"testing"
	"fmt"

)

var test_data = []byte(`{"name": "xuyifei","phone": "xuyifei","addr": "xuyifei","extern": ["x", "y", "f"]}`)

func TestParseParams(t *testing.T) {
	ParseParams(test_data)
	fmt.Println(test_data)
}