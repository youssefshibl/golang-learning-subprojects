package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type SOrder struct {
	id        int
	price     int
	createdAt string
}

type SUser struct {
	name     string
	age      int
	email    string
	password string
	orders   []SOrder
}

func (user *SUser) addUser(
	name string,
	age int,
	email string,
	password string,
) {
	user.name = name
	user.age = age
	user.email = email
	hasher := sha256.New()
	hasher.Write([]byte(password))
	user.password = hex.EncodeToString(hasher.Sum(nil))
}

func (user *SUser) addUserOrder(
	price int,
) {

	if len(user.orders) == 0 {
		now := time.Now()
		user.orders = []SOrder{{id: 0, price: price, createdAt: now.String()}}
	} else {
		id := len(user.orders) + 1
		now := time.Now()
		user.orders = append(user.orders, SOrder{id: id, price: price, createdAt: now.String()})
	}

}

func main() {

	user := SUser{}
	user.addUser("youssef", 25, "youssefshibl00@gmail.com", "1234567890")
	user.addUserOrder(25)
	user.addUserOrder(30)
	fmt.Println(user)
}
