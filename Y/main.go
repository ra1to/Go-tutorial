package main

import "fmt"

// func main() {
// 	numbers := []int{10, 20, 30}
// 	for i, n := range numbers {
// 		fmt.Println(i, n)
// 	}
// }

// スライス
// func main() {
// 	orders := make([]string, 0, 100)
// 	orders = append(orders, "ハンバーガー")
// 	orders = append(orders, "ポテト")
// 	orders = append(orders, "コーラ")
// 	fmt.Println(len(orders))
// }

// ポインタ
func main() {
	x := 10
	p := &x
	fmt.Println("pのポインタ:", p)
	fmt.Println("pのポインタの中身:", *p)
}
