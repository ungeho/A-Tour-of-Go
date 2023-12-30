package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// インターフェース自体の中にある具体的な値がnilの場合
// メソッドはnilをレシーバーとして呼び出される。
// この場合、いくつかの言語ではnullポインター例外を引き起こすが
// Goではnilをレシーバーとして呼び出されても適切に処理するメソッドを記述するのが一般的
// ただし、具体的な値としてnilを保持するインターフェースの値自体は非nilである事に注意。
func (t *T) M() {
	// Mメソッドでは、nilをレシーバーとして呼び出された時に、適切な処理をするように書かれている。
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

// 値と型を表示する。
// describe...説明する。
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
