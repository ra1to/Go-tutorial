package main

// type Student struct {
// 	name          string
// 	math, english float64
// }

// func (s Student) avg() {
// 	fmt.Println(s.name, (s.math+s.english)/2)
// }

// func main() {
// 	a001 := Student{name: "sato", math: 80, english: 70}
// 	a001.avg()
// }

// メソッドに引数を渡す
// type Student struct {
// 	name string
// }

// func (s Student) avg(math, english float64) {
// 	fmt.Println((math + english) / 2)
// }

// func main() {
// 	a001 := Student{"sato"}
// 	a001.avg(80, 70)
// }

//戻り値のあるメソッド

// type Student struct {
// 	name string
// }

// func (s Student) avg(math, english float64) float64 {
// 	return (math + english) / 2
// }

// func main() {
// 	a001 := Student{"sato"}
// 	fmt.Println(a001.avg(80, 70))
// }

//return省略

// type Student struct {
// 	name string
// }

// func (s Student) avg(math, english float64) (avgResult float64) {
// 	avgResult = (math + english) / 2
// 	return
// }

// func main() {
// 	a001 := Student{"sato"}
// 	fmt.Println(a001.avg(80, 70))
// }

//確認問題

// type User struct {
// 	name string
// }

// func (a User) cal(weight, height float64) (result float64) {
// 	result = weight / height / height * 10000
// 	return
// }

// func main() {
// 	user01 := User{"raito"}
// 	fmt.Println(user01.name, user01.cal(75, 165))
// }
