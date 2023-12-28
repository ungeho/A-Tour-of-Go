package main

import (
	"fmt"
	"math"
)

// 関数も変数として扱う事が出来る。
// 関数値(function value)を関数の引数に取る事も出来る。
// 引数がdoubleで二つの戻り値がdoubleの関数を引数に取り、doubleで返す。
func compute(fn func(float64, float64) float64) float64 {
	// 引数として受け取った関数を呼び出し、その関数の引数に3,4を渡して計算結果を返す。
	return fn(3, 4)
}

func main() {
	// 関数を定義し、変数に格納する事も出来る。
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	// 実質hypot(3, 4)
	fmt.Println(compute(hypot))
	// 実質math.pow(3,4)
	fmt.Println(compute(math.Pow))
}
