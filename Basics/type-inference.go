package main

import "fmt"

func main() {
	// :=やvar = を使って明示的な型を指定せずに変数を宣言した場合
	// 右辺の変数から型推論される。
	v := 42 //change me!
	fmt.Printf("v is of type %T\n", v)
}
