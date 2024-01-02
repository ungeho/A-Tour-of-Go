package main

import (
	"io"
	"os"
	"strings"
)

// 別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Reader。
// ROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出す
// アルファベットを一文字毎に13文字後のアルファベット(e.g. AはNに、BはOに)に置き換える
type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	// rot13Reader型の構造体、rotのrを読み出し、Readメソッドを呼び出す。
	// pにはアルファベット"Lbh penpxrq gur pbqr!"が入っている。
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if p[i] >= 'A' && p[i] < 'Z' { //A-Z
			// ASCIIコード表で'A'は10進数で65
			// (((p[i] - 65) + 13) % 26) : Aを0として、対象の文字を13文字進める。
			// 26で割った余りにする事で、Zより進んでもAから再度開始する形になる。
			// 最後に65を追加して、'A'から始まるASCIIコードに対応させる。
			p[i] = 65 + (((p[i] - 65) + 13) % 26)
		} else if p[i] >= 'a' && p[i] <= 'z' { //a-z
			// ASCIIコード表で'a'は10進数で97
			// 上記と同じ
			p[i] = 97 + (((p[i] - 97) + 13) % 26)
		}
	}
	return
}

func main() {
	// string.Readerを生成
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// ラップして文字列を変換する。
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
