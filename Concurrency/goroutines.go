package main

import (
	"fmt"
	"time"
)

// goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッド
// ランタイム....実行の部分を取り出したプログラム
// スレッド......プログラムにおいて特定の処理を行うための一貫性のある命令の流れ
// ルーチン......処理のかたまり（処理単位）

func say(s string) {
	// 0.1秒毎に5回表示する。
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// goをつけるとgoroutineで処理される。
	go say("world")
	say("hello")
}
