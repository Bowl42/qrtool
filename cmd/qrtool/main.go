package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bowl42/qrtool/internal/server"
)

func main() {
	addr := ":" + env("PORT", "8080")

	srv := &http.Server{
		Addr:              addr,
		Handler:           server.New(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("qrtool listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
