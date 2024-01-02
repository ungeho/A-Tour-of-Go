package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// 高さと幅が定義されたImage型構造体
type Image struct {
	h int
	w int
}

// ColorModel は、 color.RGBAModel を返すように
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds は、 image.Rect(0, 0, w, h) のようにして image.Rectangle を返す
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At は、ひとつの色を返す。
// 生成する画像の色の値 v を color.RGBA{v, v, 255, 255} を利用して返す
func (i Image) At(x, y int) color.Color {
	// 描画したい色の値
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 200}
	pic.ShowImage(m)
}
