package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// forループに利用する rangeは、スライスや、マップ（map）を一つずつ反復処理する為に使う。
	// スライスをrangeで繰り返す場合、rangeは反復毎に二つの変数を返す。
	// 一つ目の変数はインデックス（index）
	// 二つ目の変数はインデックスの場所の要素のコピー
	// スライスpowの最初から最後まで繰り返し
	// インデックスをiに、インデックスの場所の要素のコピーをvに格納して処理を行う。
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
