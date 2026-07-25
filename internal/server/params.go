package server

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/Bowl42/qrtool/internal/qr"
)

const (
	defaultFilename = "qrcode"
	defaultSize     = 256
	defaultLevel    = "m"
	defaultMargin   = 4
)

type params struct {
	Text     string
	Filename string
	Format   qr.Format
	Size     int
	Level    qr.Level
	Margin   int
}

type paramError struct {
	status  int
	message string
}

func (e paramError) Error() string {
	return e.message
}

func parseParams(r *http.Request) (params, error) {
	if r.Method != http.MethodGet {
		return params{}, paramError{status: http.StatusMethodNotAllowed, message: "method not allowed"}
	}

	filename, format, err := parsePath(r.URL.Path)
	if err != nil {
		return params{}, err
	}

	q := r.URL.Query()
	codeType := defaultString(q.Get("type"), "qr")
	if strings.ToLower(codeType) != "qr" {
		return params{}, paramError{status: http.StatusBadRequest, message: "unsupported type"}
	}

	text := q.Get("text")
	if text == "" {
		return params{}, paramError{status: http.StatusBadRequest, message: "missing text"}
	}

	size, err := parseInt(q.Get("size"), defaultSize, 64, 2048, "size")
	if err != nil {
		return params{}, err
	}

	margin, err := parseInt(q.Get("margin"), defaultMargin, 0, 20, "margin")
	if err != nil {
		return params{}, err
	}

	level, err := qr.ParseLevel(defaultString(q.Get("level"), defaultLevel))
	if err != nil {
		return params{}, paramError{status: http.StatusBadRequest, message: "invalid level"}
	}

	return params{
		Text:     text,
		Filename: filename,
		Format:   format,
		Size:     size,
		Level:    level,
		Margin:   margin,
	}, nil
}

func parsePath(requestPath string) (string, qr.Format, error) {
	clean := path.Clean("/" + requestPath)
	if clean == "/" {
		return defaultFilename, qr.FormatPNG, nil
	}

	if strings.Trim(clean, "/") != path.Base(clean) {
		return "", "", paramError{status: http.StatusNotFound, message: "not found"}
	}

	base := path.Base(clean)
	ext := strings.TrimPrefix(path.Ext(base), ".")
	if ext == "" {
		return "", "", paramError{status: http.StatusBadRequest, message: "unsupported format"}
	}

	format, err := qr.ParseFormat(ext)
	if err != nil {
		return "", "", paramError{status: http.StatusBadRequest, message: "unsupported format"}
	}

	name := strings.TrimSuffix(base, "."+ext)
	if name == "" {
		name = defaultFilename
	}

	return sanitizeFilename(name), format, nil
}

func parseInt(raw string, fallback, minValue, maxValue int, name string) (int, error) {
	if raw == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil || value < minValue || value > maxValue {
		return 0, paramError{
			status:  http.StatusBadRequest,
			message: fmt.Sprintf("invalid %s", name),
		}
	}

	return value, nil
}

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, `"`, "")
	name = strings.ReplaceAll(name, "\\", "")
	if name == "" || name == "." || name == ".." {
		return defaultFilename
	}
	return name
}

func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

func writeParamError(w http.ResponseWriter, err error) bool {
	var pe paramError
	if !errors.As(err, &pe) {
		return false
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(pe.status)
	_, _ = w.Write([]byte(pe.message + "\n"))
	return true
}
