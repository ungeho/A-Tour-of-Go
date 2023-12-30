package main

import "fmt"

// ゼロ個のメソッドを指定されたインターフェース型は空のインターフェースと呼ばれる。
// interface{}
// 空のインターフェースは、任意の型を保持できる。
// 全ての型は、少なくともゼロ個のメソッドを実装している。

// 空のインターフェースは、未知の型を扱うコードで使用される。
// 例えば、 fmt.Print は interface{} 型の任意の数の引数を受け取る。

func main() {
	// 空のインターフェース
	var i interface{}

	describe(i)

	// ゼロ個のメソッドが実装された整数型のインターフェース
	i = 42
	describe(i)

	// ゼロ個のメソッドが実装された文字列型のインターフェース
	i = "hello"
	describe(i)

	// 任意の型を受け入れる。
	describe(5.5)
}

// 引数に空のインターフェースを指定しているので、任意の型を引数に取る事が出来る。
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
