package qr

import (
	"bytes"
	"fmt"
	"html"
	"image"
	"image/color"
	"image/png"
	"strings"

	goqrcode "github.com/skip2/go-qrcode"
)

type Level string

const (
	LevelLow      Level = "l"
	LevelMedium   Level = "m"
	LevelQuartile Level = "q"
	LevelHigh     Level = "h"
)

type Format string

const (
	FormatPNG Format = "png"
	FormatSVG Format = "svg"
)

type Options struct {
	Text   string
	Format Format
	Level  Level
	Size   int
	Margin int
}

type Image struct {
	Format Format
	Body   []byte
}

func ParseLevel(value string) (Level, error) {
	switch strings.ToLower(value) {
	case string(LevelLow):
		return LevelLow, nil
	case string(LevelMedium):
		return LevelMedium, nil
	case string(LevelQuartile):
		return LevelQuartile, nil
	case string(LevelHigh):
		return LevelHigh, nil
	default:
		return "", fmt.Errorf("unknown qr level: %s", value)
	}
}

func ParseFormat(value string) (Format, error) {
	switch strings.ToLower(value) {
	case string(FormatPNG):
		return FormatPNG, nil
	case string(FormatSVG):
		return FormatSVG, nil
	default:
		return "", fmt.Errorf("unknown qr format: %s", value)
	}
}

func (f Format) ContentType() string {
	switch f {
	case FormatSVG:
		return "image/svg+xml; charset=utf-8"
	default:
		return "image/png"
	}
}

func Generate(opts Options) (Image, error) {
	matrix, err := bitmap(opts.Text, opts.Level)
	if err != nil {
		return Image{}, err
	}

	var body []byte
	switch opts.Format {
	case FormatPNG:
		body, err = renderPNG(matrix, opts.Size, opts.Margin)
	case FormatSVG:
		body, err = renderSVG(matrix, opts.Text, opts.Margin)
	default:
		return Image{}, fmt.Errorf("unsupported qr format: %s", opts.Format)
	}
	if err != nil {
		return Image{}, err
	}

	return Image{
		Format: opts.Format,
		Body:   body,
	}, nil
}

func renderPNG(matrix [][]bool, size, margin int) ([]byte, error) {
	modules := len(matrix) + margin*2
	if modules <= 0 {
		return nil, fmt.Errorf("invalid qr dimensions")
	}

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	black := color.RGBA{A: 255}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.Set(x, y, white)
		}
	}

	for y, row := range matrix {
		for x, dark := range row {
			if !dark {
				continue
			}
			x0 := (x + margin) * size / modules
			x1 := (x + margin + 1) * size / modules
			y0 := (y + margin) * size / modules
			y1 := (y + margin + 1) * size / modules
			for py := y0; py < y1; py++ {
				for px := x0; px < x1; px++ {
					img.Set(px, py, black)
				}
			}
		}
	}

	var out bytes.Buffer
	if err := png.Encode(&out, img); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func renderSVG(matrix [][]bool, title string, margin int) ([]byte, error) {
	size := len(matrix) + margin*2
	if size <= 0 {
		return nil, fmt.Errorf("invalid qr dimensions")
	}

	var out strings.Builder
	out.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d" shape-rendering="crispEdges">`, size, size))
	out.WriteString(`<rect width="100%" height="100%" fill="#fff"/>`)
	out.WriteString(`<path fill="#000" d="`)

	for y, row := range matrix {
		for x, dark := range row {
			if dark {
				out.WriteString(fmt.Sprintf("M%d %dh1v1h-1z", x+margin, y+margin))
			}
		}
	}

	out.WriteString(`"/>`)
	out.WriteString(`<title>`)
	out.WriteString(html.EscapeString(title))
	out.WriteString(`</title></svg>`)
	return []byte(out.String()), nil
}

func bitmap(text string, level Level) ([][]bool, error) {
	code, err := goqrcode.New(text, recoveryLevel(level))
	if err != nil {
		return nil, err
	}
	code.DisableBorder = true
	return code.Bitmap(), nil
}

func recoveryLevel(level Level) goqrcode.RecoveryLevel {
	switch level {
	case LevelLow:
		return goqrcode.Low
	case LevelQuartile:
		return goqrcode.High
	case LevelHigh:
		return goqrcode.Highest
	default:
		return goqrcode.Medium
	}
}
