package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在の日時を取得する。
	t := time.Now()

	// 取得した現在の時刻の部分を比較し、対応したケースの出力をする。
	// 条件のないswitchは switch trueと同様
	// 長くなりがちな"if-then-else"のつながりをシンプルに表現出来る。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
