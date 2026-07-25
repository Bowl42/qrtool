package server

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"mime"
	"net/http"

	"github.com/Bowl42/qrtool/internal/qr"
)

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("ok\n"))
}

func handleQR(w http.ResponseWriter, r *http.Request) {
	if shouldShowHome(r) {
		handleHome(w, r)
		return
	}

	p, err := parseParams(r)
	if err != nil {
		if writeParamError(w, err) {
			return
		}
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	etag := makeETag(p)
	if r.Header.Get("If-None-Match") == etag {
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	image, err := qr.Generate(qr.Options{
		Text:   p.Text,
		Format: p.Format,
		Level:  p.Level,
		Size:   p.Size,
		Margin: p.Margin,
	})
	if err != nil {
		http.Error(w, "failed to generate qr", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", image.Format.ContentType())
	w.Header().Set("Content-Disposition", contentDisposition(p))
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Header().Set("ETag", etag)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(image.Body)
}

func contentDisposition(p params) string {
	value := mime.FormatMediaType("inline", map[string]string{
		"filename": p.Filename + "." + string(p.Format),
	})
	if value == "" {
		return "inline"
	}
	return value
}

func makeETag(p params) string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("qr|%s|%s|%d|%d|%s", p.Format, p.Text, p.Size, p.Margin, p.Level)))
	return `"` + hex.EncodeToString(hash[:]) + `"`
}
