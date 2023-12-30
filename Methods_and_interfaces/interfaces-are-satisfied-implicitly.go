package main

import "fmt"

// 型にメソッドを実装しておくことで、インターフェースを実装する。
type I interface {
	M()
}

// 構造体
type T struct {
	S string
}

// インターフェースに取りまとめられたメソッド
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// インターフェースを満たした型で変数iを定義して
	// インターフェースに取りまとめられたメソッドの引数の構造体を代入する。
	var i I = T{"hello"}

	// メソッドを実行する。
	i.M()
}
