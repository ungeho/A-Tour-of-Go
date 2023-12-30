package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// 通常の関数の場合、変数の引数を取る事を指定している為、もしここで&vを渡すと、コンパイルエラー
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// メソッドの場合、変数の引数を取る事を指定していて、アドレスを渡したとしても、エラーにならず
	// 変数、または、ポインタのいずれかのレシーバーとして機能してくれる。
	// その為、中身を書き換えたい時はポインタ、値を使いたいだけの時は変数、だけを意識して使用出来る。
	// この場合、p.Abs() は (*p).Abs() として解釈される。
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
