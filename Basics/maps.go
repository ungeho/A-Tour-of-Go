package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// make関数は指定された型のマップを初期化して、使用可能な状態で返す。
	m = make(map[string]Vertex)
	// mapはキーと値を関連付ける。（マップする）
	// "Bell Labs"という名前のキーと構造値のフィールドに入れた値を関連付ける。
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// マップのゼロ値はnil。nilマップはキーを持たず、キーを追加する事も出来ない。
}
