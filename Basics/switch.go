package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// Goのswitchはcaseの最後に必要なbreakが自動的に提供される為
	// 選択されたcaseのみを実行し、続く全てのcaseは実行されない。
	// また、Goのswitchのcaseは定数である必要はなく、関係する値も整数である必要はない。

	// 変数osにruntime.GOOSによって取得したOSの名前を格納し、osとcaseが一致したものを出力する。
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows....
		fmt.Printf("%s.\n", os)
	}
}
