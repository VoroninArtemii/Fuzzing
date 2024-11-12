package fuzz

import (
	"testing"
)

func FuzzNothing(f *testing.F) {
	f.Add([]byte("\xa2\x4aapiVersion\x41x\x44kind\x41y"))
	f.Fuzz(func(t *testing.T, in []byte) {})
}