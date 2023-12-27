package main

// フォーマットパッケージとtimeパッケージをimport
import (
	"fmt"
	"time"
)

func main() {
	// 文字列を表示。最後に改行
	fmt.Println("Welcome to the Playground")
	// 文字列と現在の時間を並べて表示。最後に改行
	fmt.Println("The time is", time.Now())
}
