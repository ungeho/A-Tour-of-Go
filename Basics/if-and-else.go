package main

import (
	"fmt"
	"math"
)

// 倍精度浮動小数点数でx,n,limの引数を取り、倍精度で返す。
func pow(x, n, lim float64) float64 {
	//ifステートメント内で宣言された変数vは、elseブロック内でも使用する事が出来る。
	// x^nがlim未満ならx^nを返す
	// そうでないなら、x^n > vである事を出力し、limを返す。
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// 3^2 < 10 なので9
	// 3^3 >= 20なのでlimである20が出力される。
	// また、Printlnの前にpowによる出力が先に入るので、2^3 >= 20 の出力があった後
	// 3^2 20 の出力が起こる。
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
