package main

import "fmt"

func main() {

	if age := 5; age >= 5 && age < 10 {
		fmt.Println("5歳以上10歳未満")
	} else {
		fmt.Println("それ以外")
	}
}
