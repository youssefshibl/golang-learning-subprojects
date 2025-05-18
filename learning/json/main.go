package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type myInt struct {
		IntValue  int     `json:"joo"`
		TestEmpyt *string `json:"test,omitempty"`
	}

	// youssef := "youssef"
	data := myInt{IntValue: 1234, TestEmpyt: nil}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
