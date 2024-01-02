package main

// io パッケージは
// データストリームを読むことを表現する io.Reader インタフェースを規定している。
// Goの標準ライブラリには
// ファイル、ネットワーク接続、圧縮、暗号化などで
// このインタフェースの 多くの実装 があります。

// io.Reader インタフェースは Read メソッドを持ちます:
// func (T) Read(b []byte) (n int, err error)
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// strings.Readerを作成
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		// Readは、データを与えられたバイトスライスへ入れ
		// 入れたバイトのサイズとエラーの値を返す。
		// ストリームの終端は、io.EOFのエラーで返す。
		n, err := r.Read(b)
		// 8バイト毎に読み出す。
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// 終端でループを抜ける
		if err == io.EOF {
			break
		}
	}
}
