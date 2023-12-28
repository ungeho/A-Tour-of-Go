package main

import (
	"fmt"
	"math"
)

// 倍精度(double)の引数xを受け取り、文字列で返す。
func sqrt(x float64) string {
	//xが0未満（負の値）であるとき、-x（絶対値x）を引数にsqrtを再起呼び出しを行い
	// 結果として返ってくるsqrt(絶対値x)の計算結果の文字列と複素数を表す"i"の文字列を連結して返す。
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// Spritnt...与えられた値や変数（オペランド）にデフォルトの書式を使用して書式を設定し、結果の文字列を返す。
	//			複数与えられたオペランドがどちらも文字列でない場合は、オペランド間にスペースが追加される。
	// xが0以上ならx^(1/2)を計算し、結果を文字列で返す。
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
