package qrcode

import (
	"github.com/skip2/go-qrcode"
	"image/color"
)

var (
	White = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	Black = color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 255,
	}
	Yellow = color.RGBA{
		R: 255,
		G: 255,
		B: 0,
		A: 255,
	}
	LightBlue = color.RGBA{
		R: 0,
		G: 255,
		B: 255,
		A: 255,
	}
	Blue = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
	Red = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	Pink = color.RGBA{
		R: 255,
		G: 0,
		B: 255,
		A: 255,
	}
	LightPurple = color.RGBA{
		R: 127,
		G: 0,
		B: 255,
		A: 255,
	}
	Green = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
	Orange = color.RGBA{
		R: 255,
		G: 127,
		B: 0,
		A: 255,
	}
)

// CreateQRCode 创建二维码
// 注意RGBA的A 不是百分制 而是 0-255
func CreateQRCode(content string, bkColor, picColor color.RGBA, size int) ([]byte, error) {
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	qr.BackgroundColor = bkColor
	qr.ForegroundColor = picColor
	png, err := qr.PNG(size)
	if err != nil {
		return nil, err
	}
	return png, nil
}
