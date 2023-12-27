package main

import "fmt"

func main() {
	// 初期値を与えずに宣言すると、ゼロ値(zero value)が与えられる。
	// 数値(intやfloat)は0
	// boolはfalse
	// stringは空文字列(empty string)""
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
