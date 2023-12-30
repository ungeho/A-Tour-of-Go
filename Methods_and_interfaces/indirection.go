package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// メソッドだと、呼び出し時に変数、またはポインタのいずれかのレシーバとしてとる事が出来る。
	// (&v).Scale(2)と解釈して実行されている。
	v.Scale(2)
	// ScaleFunc(関数)だと、指し示す先（ポインタ）を渡す必要がある。
	// 特に、下記の1行で&を忘れるとコンパイルエラー
	ScaleFunc(&v, 10)

	// 予め指し示す先をpに入れておく。
	p := &Vertex{4, 3}
	// メソッドでは変数とポインタ、どちらを渡しても良い感じに処理してくれる。
	p.Scale(3)
	// pは既に指し示す先を渡してるのでOK
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
