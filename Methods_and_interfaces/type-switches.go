package main

import "fmt"

// 空のインターフェースを引数に取り、どんな型も引数として受け入れる。
func do(i interface{}) {
	// 型switch...いくつかの型アサーションを直列に使用できる。
	// 型switchは通常のswitchと似ているが、型switchのcaseは型（値ではない）を指定し、
	// それらの値は指定されたインターフェースの値が保持する値の型と比較される。

	// ここで、型アサーションi.(T)と同じ構文を持つが、特定の型Tはキーワードtypeに置き換えられる。
	switch v := i.(type) {
	case int:
		// 型がint型の場合
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// 型が文字列型の場合
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		// (int,文字列)以外の型の場合
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
