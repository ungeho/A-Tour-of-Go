package main

import "fmt"

// もっともよく使われているinterfaceの一つが
// fmtパッケージに定義されているStringer
// type Stringer interface {
// 	String() string
// }

// 人間の構造体 名前と年齢
type Person struct {
	Name string
	Age  int
}

// Stringerインターフェースはstring（文字列）として表現出来る型
// fmtパッケージと、その他の多くのパッケージでは、変数を文字列で出力する為にこのインターフェーズがある事を確認する。
// String()のメソッドを作成することで
// fmtパッケージにPersonインターフェースを渡した時の表示方法が変わる。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	// 通常の場合
	// {Arthur Dent 42} {Zaphod Beeblebrox 9001}
	// のように構造体の形式で表示されるが
	// String()メソッドを定義したことで
	// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
	// と表示される。
	fmt.Println(a, z)
}
