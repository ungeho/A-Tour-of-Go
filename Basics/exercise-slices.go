package main

// ローカル環境ではimport出来なかったので、A tour of Goのブラウザ上で実行する。
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 容量=dxの二次元スライスを作成
	image := make([][]uint8, dy)

	// imageの最初から最後まで繰り返す(dx)
	// 必要なのはindexのみな為、valueは省略
	// indexはy座標として考える。
	for y := range image {
		// スライスの中にスライスを割り当てる。
		image[y] = make([]uint8, dx)

		// 今の処理中のy座標で定義したx座標方向全てに対して
		// 二次元配列の横のindexに対応する値をxに格納し
		// xとyを使った計算結果を格納する。
		for x := range image[y] {
			// xとyはint型、imageはunsigned int(8bit)なので、計算結果をuint8にキャストして格納する。
			// image[y][x] = uint8((x + y) / 2)
			// image[y][x] = uint8(x * y)
			image[y][x] = uint8(x ^ y)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
