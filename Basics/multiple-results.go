package main

import "fmt"

// 複数の戻り値を返す事が出来る。
// ()で囲って複数の戻り値の型を指定する
// その為、簡単にswap出来る。
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// ,で区切って複数の戻り値を代入する。
	// 代入は=ではなく:=
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
