package main

import "fmt"

// 関数は0個以上の引数を取る事が出来る。
// また、引数、関数ともに型は変数名の後ろに書く事に注意
// 詳しくは:https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
