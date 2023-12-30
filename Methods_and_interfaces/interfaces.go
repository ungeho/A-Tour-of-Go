package main

import (
	"fmt"
	"math"
)

// interface(インターフェース)型は、メソッドのシグニチャの集まりで定義する。
// シグニチャ...（署名,サイン）メソッドを識別するもの。特にメソッド名、引数の型、引数の数の3要素
// メソッドの集まりを実装した値を、interface型の変数へ持たせることが出来る。
type Abser interface {
	// 倍精度で返すAbsメソッドがひとまとめにされている。
	// MyFloatを引数に取るAbsメソッド
	// (*Vertex)を引数に取るAbsメソッド
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// MyFloatはAbserを実装（implemetnts）
	a = f
	// (*Vertex)はAbserを実装（implemetnts）
	a = &v

	// vはVertex
	// Abserが実装されていないので以下の1行はエラーになる。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloatを引数に取るAbsメソッド
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// (*Vertex)を引数に取るAbsメソッド
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
