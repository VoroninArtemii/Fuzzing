package parsers

import (
	"testing"
	dockerref "github.com/distribution/reference"
)

func FuzzParseImageName(f *testing.F) {
	f.Add("url")
	f.Add(":")
	f.Add("/")
	f.Add("@")
	f.Add("sha256")
	f.Add("repo")
	f.Fuzz(func(t *testing.T, input string) {
		_, err_1 := dockerref.ParseNormalizedNamed(input)
		if err_1 != nil {
			t.Skip()
		}
		_, _, _, err_2 := ParseImageName(input)
		if (err_2 != nil){
				t.Errorf(input, err_2)
		}
	})
}
