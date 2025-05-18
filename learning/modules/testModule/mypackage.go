package testModule

import "fmt"

func returnName() string {
	return "Youssef"
}

func PrintHello() {
	fmt.Println("Hello,", returnName())
}
