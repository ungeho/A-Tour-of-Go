package main

import (
	"fmt"
	"math"
)

// 引数x,n,limは全て倍精度。 返す値も倍精度
func pow(x, n, lim float64) float64 {
	// ifステートメントはforのように条件を評価する前に
	// 評価する為の簡単なステートメントを記述する事が出来る。
	// ここで宣言された変数は、ifスコープ内だけで有効。
	// x^nを計算し、vに代入。 その後、v<limを評価し、lim未満であればx^nを返す。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// x^nがlim以上の値だった場合、limを返す。
	return lim
}

func main() {
	// 複数行に分けて記述する場合、最後の引数の後にも,が必要
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
