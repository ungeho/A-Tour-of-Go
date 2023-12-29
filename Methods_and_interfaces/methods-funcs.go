package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッドは、レシーバ引数（funcとAbsの間に引数）を伴う関数。
// 通常の書き方で関数を定義することで
// methods.go の例から機能を変えずに通常の関数として記述できる。
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
