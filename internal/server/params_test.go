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
	req := httptest.NewRequest("GET", "/demo.svg?text=hello&size=512&margin=2&level=h", nil)

	p, err := parseParams(req)
	if err != nil {
		t.Fatalf("parseParams() error = %v", err)
	}

	if p.Filename != "demo" || p.Format != "svg" || p.Size != 512 || p.Margin != 2 || p.Level != "h" {
		t.Fatalf("unexpected params: %+v", p)
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
