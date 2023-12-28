package main

import "fmt"

func main() {
	// 容量10の空の配列を示すスライスを作成
	pow := make([]int, 10)

	// インデックスと値の片方のみが必要な場合はアンダーバー"_"へ捨てる事が出来る。

	// i , _ と同じ、インデックスのみが必要な場合、二つ目の値（value）を省略する。
	for i := range pow {
		// 2**i
		pow[i] = 1 << uint(i)
	}

	// スライスpowの値(value)だけが欲しいのでインデックス(index)をアンダーバー"_"で捨てる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
