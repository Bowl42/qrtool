package server

import (
	"net/http/httptest"
	"testing"
)

func TestParseParamsDefaults(t *testing.T) {
	req := httptest.NewRequest("GET", "/?text=hello", nil)

	p, err := parseParams(req)
	if err != nil {
		t.Fatalf("parseParams() error = %v", err)
	}

	if p.Filename != "qrcode" || p.Format != "png" || p.Text != "hello" || p.Size != 256 || p.Margin != 4 {
		t.Fatalf("unexpected params: %+v", p)
	}
}

func TestParseParamsAcceptsExplicitQRType(t *testing.T) {
	req := httptest.NewRequest("GET", "/?type=qr&text=hello", nil)

	if _, err := parseParams(req); err != nil {
		t.Fatalf("parseParams() error = %v", err)
	}
}

func TestParseParamsPath(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.svg?text=hello&size=512&margin=2&level=h&fg=123456&bg=abcdef", nil)

	p, err := parseParams(req)
	if err != nil {
		t.Fatalf("parseParams() error = %v", err)
	}

	if p.Filename != "demo" || p.Format != "svg" || p.Size != 512 || p.Margin != 2 || p.Level != "h" {
		t.Fatalf("unexpected params: %+v", p)
	}
	if p.Foreground.R != 0x12 || p.Foreground.G != 0x34 || p.Foreground.B != 0x56 {
		t.Fatalf("unexpected foreground: %+v", p.Foreground)
	}
	if p.Background.R != 0xab || p.Background.G != 0xcd || p.Background.B != 0xef {
		t.Fatalf("unexpected background: %+v", p.Background)
	}
}

func TestParseParamsRequiresText(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.png", nil)

	_, err := parseParams(req)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseParamsRejectsUnsupportedType(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.png?type=code128&text=hello", nil)

	_, err := parseParams(req)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseParamsRejectsUnsupportedFormat(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.jpg?text=hello", nil)

	_, err := parseParams(req)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestParseParamsRejectsInvalidColor(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.png?text=hello&fg=wat", nil)

	_, err := parseParams(req)
	if err == nil {
		t.Fatal("expected error")
	}
}
