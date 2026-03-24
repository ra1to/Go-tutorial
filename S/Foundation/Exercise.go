package main

import "fmt"

func main() {
	// Q1 1.11をint型に変換して出力
	// f := 1.11
	// fmt.Println(int(f))

	// 正解
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	// Q2 5と6が出力される

	// Q3

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
