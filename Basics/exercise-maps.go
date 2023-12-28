package main

// ローカルではimport出来ないので、ブラウザ上で実行する。
import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// 単語をキーに、出現回数をカウントしたいので
	// 文字列のキー値を持つ、要素がint型のマップを定義する。
	wc := make(map[string]int)

	// strings.Fieldsは与えられた文字列sを元に、一つ以上の空白で分割した文字列のスライスを返す。
	// indexは必要なく、valueである単語(word)が欲しいのでindexは_で捨てる。
	for _, w := range strings.Fields(s) {
		// 単語をキーとして、値を一つ増やす。（intのゼロ値は0なので、カウントするように増えていく。）
		wc[w]++
	}

	return wc
}

func main() {
	wc.Test(WordCount)
}
