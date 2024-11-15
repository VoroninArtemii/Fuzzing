package fuzz

import (
	"testing"
)

func FuzzParseImageName(f *testing.F) {
	f.Add("url:5000/repo@sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	f.Fuzz(func(t *testing.T, input string) {})
}
