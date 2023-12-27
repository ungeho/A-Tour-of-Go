package main

import "fmt"

// var宣言では、変数毎に初期化子(initializer)を与えられる。
// int型の変数i,jを初期値1,2で宣言
var i, j int = 1, 2

func main() {
	// 初期化子が与えられていると、型の宣言を省略する事が出来る。
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
