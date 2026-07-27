package qr

import (
	"image/color"
	"strings"
	"testing"
)

func TestGeneratePNG(t *testing.T) {
	image, err := Generate(Options{
		Text:   "hello",
		Format: FormatPNG,
		Level:  LevelMedium,
		Size:   256,
		Margin: 4,
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if image.Format != FormatPNG {
		t.Fatalf("Format = %q", image.Format)
	}
	if len(image.Body) < 8 || string(image.Body[:8]) != "\x89PNG\r\n\x1a\n" {
		t.Fatal("output is not a PNG")
	}
}

func TestGenerateSVG(t *testing.T) {
	image, err := Generate(Options{
		Text:   "hello",
		Format: FormatSVG,
		Level:  LevelMedium,
		Size:   256,
		Margin: 4,
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if image.Format != FormatSVG {
		t.Fatalf("Format = %q", image.Format)
	}
	if !strings.HasPrefix(string(image.Body), `<svg `) {
		t.Fatal("output is not an SVG")
	}
}

func TestGenerateSVGUsesColors(t *testing.T) {
	image, err := Generate(Options{
		Text:       "hello",
		Format:     FormatSVG,
		Level:      LevelMedium,
		Size:       256,
		Margin:     4,
		Foreground: colorRGBA(0x12, 0x34, 0x56),
		Background: colorRGBA(0xab, 0xcd, 0xef),
	})
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	body := string(image.Body)
	if !strings.Contains(body, `fill="#abcdef"`) {
		t.Fatalf("missing background color: %s", body)
	}
	if !strings.Contains(body, `fill="#123456"`) {
		t.Fatalf("missing foreground color: %s", body)
	}
}

func TestParseFormat(t *testing.T) {
	if _, err := ParseFormat("svg"); err != nil {
		t.Fatalf("ParseFormat() error = %v", err)
	}
	if _, err := ParseFormat("jpg"); err == nil {
		t.Fatal("expected error")
	}
}

func colorRGBA(r, g, b uint8) color.RGBA {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func TestParseLevel(t *testing.T) {
	if _, err := ParseLevel("h"); err != nil {
		t.Fatalf("ParseLevel() error = %v", err)
	}
	if _, err := ParseLevel("x"); err == nil {
		t.Fatal("expected error")
	}
}
