package main

import "fmt"

func main() {
	// deferステートメント...呼び出し元の関数(main)の終わり(return)まで遅延させる
	defer fmt.Println("world")

	// 先にこちらの処理が実行される
	fmt.Println("hello")

	// 関数の終わりに、deferで遅延された処理が実行される。
}
