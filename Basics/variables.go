package main

import "fmt"

// varは変数（variable）を宣言する。
// 関数の引数リストと同様に、複数の変数の最後に型をかくことで、変数のリストを宣言出来る。
// bool型で三つの変数を宣言
var c, python, java bool

func main() {
	// int型で変数iを宣言
	var i int
	fmt.Println(i, c, python, java)
}
