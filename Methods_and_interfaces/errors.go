package main

import (
	"fmt"
	"time"
)

// Goのプログラムは、エラーの状態を error 値で表現する。
// error 型は fmt.Stringer に似た組み込みインターフェース
// type error interface {
// 		Error() string
// }
// fmt.Stringer と同様に、 fmt パッケージは、変数を文字列で出力する際に error インタフェースを確認する。

// MyError型の構造体を生成
type MyError struct {
	// Time型のWhen
	When time.Time
	// 文字列型のWhat
	What string
}

// 組み込みのErrorインターフェース
func (e *MyError) Error() string {
	// 時間と文字列を返す
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// error型の変数を返す関数。
func run() error {
	// エラーの発生時間とエラーメッセージを返す。
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// 関数は error 型の変数を返す。
// そして、呼び出し元はエラーが nil かどうかを確認することでエラーをハンドル(取り扱い)する。
// i, err := strconv.Atoi("42")
// if err != nil {
//     fmt.Printf("couldn't convert number: %v\n", err)
//     return
// }
// fmt.Println("Converted integer:", i)
// nilの error は成功したことを示し、nilでない errorは失敗したことを示す。

func main() {
	// エラー型を返す関数の返す値がnilでない場合
	// エラー発生として、エラーメッセージを出力する。
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
