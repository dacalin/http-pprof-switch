package pprof

import (
	"os"
	"strings"
)

func enabled() bool {
	if os.Getenv("ENABLE_HTTP_PPROF") == "1" {
		return true
	}

	if strings.ToLower(os.Getenv("ENABLE_HTTP_PPROF")) == "true" {
		return true
	}
	return false
}

func host() string {
	h := os.Getenv("HTTP_PPROF_HOST")

	if h == "" {
		h = "0.0.0.0"
	}

	return h
}

func port() string {
	p := os.Getenv("HTTP_PPROF_PORT")

	if p == "" {
		p = "1000"
	}

	return p
}
