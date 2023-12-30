package main

import (
	"fmt"
	"math"
)

// 任意の型（type）宣言
type MyFloat float64

// typeによる任意の型や、組み込み型にもメソッドを宣言出来る。
// また、レシーバを伴うメソッドの宣言は、レシーバ型が同じパッケージ内にある必要がある。
// 他のパッケージに定義している型に対して、レシーバに伴うメソッドを定義できない。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// -math.Sqrt2をMyFloat型にキャストし、それを変数fに初期化して定義
	f := MyFloat(-math.Sqrt2)
	// MyFloat型のメソッドであるAbsメソッドの結果を出力
	fmt.Println(f.Abs())
}
