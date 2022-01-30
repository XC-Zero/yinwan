package qrcode

import (
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"image/color"
	"strconv"
	"strings"
)

// 这是一坨定义好的颜色，可以往里面加颜色，看你们喜欢
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
// 注意:此处的RGBA的A 不是百分制 而是 0-255
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

// TransferColor 十六进制转 RGBA颜色
func TransferColor(colorStr string, transparency float64) (rgba color.RGBA, err error) {
	colorStr = strings.ReplaceAll(colorStr, "#", "")
	if len(colorStr) != 6 {
		err = errors.New("Invalid color string !")
		return
	}
	R, err := strconv.ParseUint(colorStr[:2], 16, 32)
	if err != nil {
		return
	}
	G, err := strconv.ParseUint(colorStr[2:4], 16, 32)
	if err != nil {
		return
	}
	B, err := strconv.ParseUint(colorStr[4:6], 16, 32)
	if err != nil {
		return
	}
	return color.RGBA{
		R: uint8(R),
		G: uint8(G),
		B: uint8(B),
		A: uint8(transparency * 255),
	}, nil
}
