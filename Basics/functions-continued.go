package main

import "fmt"

// 関数の引数が二つ以上同じ型である場合、最後の型を残して省略して記述できる。
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
