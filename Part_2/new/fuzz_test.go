package fuzz

import (
	"testing"
	"k8s.io/kubernetes/pkg/util/parsers"
	dockerref "github.com/distribution/reference"
)

func FuzzParseImageName(f *testing.F) {
	f.Add("url:5000/repo@sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	f.Fuzz(func(t *testing.T, input string) {
		_, err_1 := dockerref.ParseNormalizedNamed(input)
		if err_1 != nil {
			t.Skip()
		}
		_, _, _, err_2 := parsers.ParseImageName(input)
		if (err_2 != nil){
				t.Errorf(input, err_2)
		}
	})
}
