package main

import "fmt"

func main() {
	// 文字列型の2個の配列（array）
	// 型推論で配列を定義する場合
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	// Hello World
	fmt.Println(a[0], a[1])
	// [Hello World]
	fmt.Println(a)

	// :=で配列を定義する場合
	// 初期化子も{}で囲み、カンマで区切って入れられる。
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// [2 3 5 7 11 13]
	fmt.Println(primes)
}
