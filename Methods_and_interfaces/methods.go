package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Goにはクラスがないが、型にメソッド(method)を定義できる。
// メソッドは特別なレシーバ(receiver)引数を関数に取る。
// レシーバは、funcキーワードとメソッド名の間に自身の引数リストで表現する。
// 以下は、Absメソッドはvという名前のVertex型のレシーバを持つ事を意味する。

// funcと関数名（この場合はメソッド名）Abs()の間に引数リストで表現。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
