package main

import "fmt"

func main() {
	// 6個の素数を配列で初期化
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 配列は固定長。スライスは可変長
	// 型[]T は型Tのスライスを表す。
	// a[low : high]とインデックスの境界を指定すると
	// low <=  > high の半開区間が指定される。
	// 下記の指定の場合は、配列primeの要素番号が1から3までのスライスを作成する。
	var s []int = primes[1:4]
	fmt.Println(s)
}
