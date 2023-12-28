package main

import "fmt"

func main() {
	// 要素が四つ分の文字列の配列を作成
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	// 配列の中身を全て出力
	fmt.Println(names)

	//nameの要素番号0,1のスライスaを作成
	a := names[0:2]
	// nameの要素番号1,2のスライスbを作成
	b := names[1:3]
	// 配列の0,1と1,2から得られたスライスを出力
	fmt.Println(a, b)

	// スライスは配列への参照である為、スライス自体はデータを格納していない。
	// スライスは、元の配列の部分列を指し示しているので、スライスの内容を上書きすると
	// 元の配列の該当部分も書き換わる。
	b[0] = "XXX" //b[0]にはnames[1]が対応している為、names[1]が書き換わる
	// names[1]が書き換わったので、names[1]を指し示すa[2]も呼び出された時に、書き換わった後の値を出力する。
	fmt.Println(a, b)
	// names[1]が書き換わったことが確認出来る。
	fmt.Println(names)
}
