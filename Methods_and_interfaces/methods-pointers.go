package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Absメソッド
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバによるメソッド宣言
// ポインタレシーバにすると、Vertexの変数にアクセスし、中の変数の値を書き換える事が出来る。
// レシーバ自身を更新することが多いため、変数レシーバよりもポインタレシーバの方が一般的。
// Scaleでは、与えられた引数を元に、レシーバの値を拡大、縮小する。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
