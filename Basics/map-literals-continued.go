package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// リテラルの要素のVertexから推定出来る為
// mapに渡すトップレベルの型は、その型名を省略する事が出来る。
var m = map[string]Vertex{
	// ここでの型名を省略出来る。
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
