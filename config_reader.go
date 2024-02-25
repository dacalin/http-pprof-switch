package pprof

import (
	"os"
	"strings"
)

func enabled() bool {
	if os.Getenv("HTTP_PPROF_SWITCH_ENABLE") == "1" {
		return true
	}

	if strings.ToLower(os.Getenv("HTTP_PPROF_SWITCH_ENABLE")) == "true" {
		return true
	}
	return false
}

func host() string {
	h := os.Getenv("HTTP_PPROF_SWITCH_HOST")

	if h == "" {
		h = "0.0.0.0"
	}

	return h
}

func port() string {
	p := os.Getenv("HTTP_PPROF_SWITCH_PORT")

	if p == "" {
		p = "1000"
	}

	return p
}
