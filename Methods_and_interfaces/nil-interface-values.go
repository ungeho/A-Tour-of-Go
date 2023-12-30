package main

import "fmt"

// nilインターフェースの値は、値も具体的な型も保持しない。
// 呼び出す具体的なメソッドを示す型が、インターフェースのタプル内に存在しない為
// nilインターフェースのメソッドを呼び出すと、ランタイムエラーになる。
type I interface {
	M()
}

func main() {
	// nilインターフェース
	var i I
	// nilインターフェースは値も具体的な型も保持していない為
	// 値も型もnilになる。
	describe(i)
	// 呼び出すとランタイムエラー
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
