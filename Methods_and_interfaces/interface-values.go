package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 構造体Tのポインタを引数とするメソッド
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// 倍精度Fを引数とするメソッド
func (f F) M() {
	fmt.Println(f)
}

func main() {
	//インターフェースの値は、値と具体的な型のタプルのように考える事が出来る。
	// (value, type)
	// インターフェースの値は、特定の基底になる具体的な型の値を保持する。
	// インターフェースの値のメソッドを呼び出すと、その基底型の同じ名前のメソッドが実行される。

	// インターフェースが実装された型で変数iを定義
	var i I

	// 構造体Tの型で、値"Hello"を入れる。
	i = &T{"Hello"}
	// 関数describeで型と値を表示
	// 値：&{"Hello"}と型：構造体(*T)が表示される。
	describe(i)
	// 構造体Tのポインタを引数とするメソッドが呼び出され、値 Hello が表示される。
	i.M()

	// 倍精度Fの型で、値πを入れる。
	i = F(math.Pi)
	// 値πと型Fが表示される。
	describe(i)
	// 倍精度Fの型を引数とするメソッドが呼び出され、値πが表示される。
	i.M()
}

// 型と値を表示する関数。
// 引数はインターフェースが実装された型、I
func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)
}
