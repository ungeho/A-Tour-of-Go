package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	// 現在の時刻から曜日を取得
	// 恐らく日曜日を0として、月曜日なら1、火曜日なら2のように値が入る。
	today := time.Now().Weekday()
	// 土曜日なので6が入る。todayから二日先まで確認し、対応した文字列を出力する。
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
