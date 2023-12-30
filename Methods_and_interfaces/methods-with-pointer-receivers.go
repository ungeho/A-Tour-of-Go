package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// ポインタレシーバを使う理由は二つある。

// 一つはメソッドがレシーバが指す先の変数の中身を変更する為。
// Scaleはレシーバが指す先の変数の値を変更している。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// もう一つはメソッドの呼び出し毎に変数のコピーを避ける為。
// レシーバが大きな構造体である場合に特に有効。
// Absは変数の中身を書き換える必要はないが、呼び出される度に構造体のコピーを作らない。
// 一般的には、値レシーバとポインタレシーバのどちらかで全てのメソッドを与え、混在させるべきではない。
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling : %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling : %+v, Abs: %v\n", v, v.Abs())
}
