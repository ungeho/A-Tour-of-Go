package main

import "fmt"

func fibonacci() func() int {
	fn, fm := 0, 1
	// 中の関数が呼び出され、変数fnとfmの値は保持されるので呼び出す度にフィボナッチ数列の項が進む
	return func() int {
		f := fn
		fn, fm = fm, fn+fm
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
