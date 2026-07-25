package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rec := httptest.NewRecorder()

	New().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	if rec.Body.String() != "ok\n" {
		t.Fatalf("body = %q", rec.Body.String())
	}
}

func TestHandleQRPNG(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.png?text=hello", nil)
	rec := httptest.NewRecorder()

	New().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Type"); got != "image/png" {
		t.Fatalf("Content-Type = %q", got)
	}
	if got := rec.Header().Get("Content-Disposition"); got != `inline; filename=demo.png` {
		t.Fatalf("Content-Disposition = %q", got)
	}
	if !strings.HasPrefix(rec.Body.String(), "\x89PNG") {
		t.Fatal("response is not a PNG")
	}
}

func TestHandleQRSVG(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo.svg?text=hello", nil)
	rec := httptest.NewRecorder()

	New().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Type"); got != "image/svg+xml; charset=utf-8" {
		t.Fatalf("Content-Type = %q", got)
	}
	if !strings.HasPrefix(rec.Body.String(), `<svg `) {
		t.Fatal("response is not an SVG")
	}
}

func TestHandleQRContentDispositionEscapesFilename(t *testing.T) {
	req := httptest.NewRequest("GET", "/demo%20code.svg?text=hello", nil)
	rec := httptest.NewRecorder()

	New().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Content-Disposition"); got != `inline; filename="demo code.svg"` {
		t.Fatalf("Content-Disposition = %q", got)
	}
}

func TestHandleQRETag(t *testing.T) {
	req := httptest.NewRequest("GET", "/a.png?text=hello", nil)
	rec := httptest.NewRecorder()
	New().ServeHTTP(rec, req)

	etag := rec.Header().Get("ETag")
	if etag == "" {
		t.Fatal("missing ETag")
	}

	req = httptest.NewRequest("GET", "/a.png?text=hello", nil)
	req.Header.Set("If-None-Match", etag)
	rec = httptest.NewRecorder()
	New().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotModified {
		t.Fatalf("status = %d", rec.Code)
	}
}
