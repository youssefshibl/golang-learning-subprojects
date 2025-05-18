// Parsing JSON Using a Struct
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	userJson := "{\"Name\": \"youssef\"}"

	type User struct {
		Name string
	}
	var user1 User
	err := json.Unmarshal([]byte(userJson), &user1)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}
	fmt.Println(user1)

}
