package parsers

import (
	"math/rand"
	"strings"
	"testing"
	"time"
	dockerref "github.com/distribution/reference"
)

func mutateInput(input string) string {
	rand.Seed(time.Now().UnixNano())
	mutations := []func(string) string{
		func(s string) string { return s + ":latest" },
		func(s string) string { return s + "@sha256:deadbeef" },
		func(s string) string { return strings.ReplaceAll(s, ":", "//") },
		func(s string) string { return "docker.io/library/" + s },
		func(s string) string { return strings.ToUpper(s) },
		func(s string) string { return s + "/" },
		func(s string) string { return s + "@sha512:" + strings.Repeat("a", 128) },
		func(s string) string { return s + ":t@gged" },
		func(s string) string { return strings.Repeat("a/", 10) + s },
	}
	return mutations[rand.Intn(len(mutations))](input)
}

func FuzzParseImageName(f *testing.F) {
	f.Add("url")
	f.Add(":")
	f.Add("/")
	f.Add("@")
	f.Add("sha256")
	f.Add("repo")
	f.Fuzz(func(t *testing.T, input string) {
		mutated := mutateInput(input)
		_, err_1 := dockerref.ParseNormalizedNamed(mutated)
		if err_1 != nil {
			t.Skip()
		}
		_, _, _, err_2 := ParseImageName(mutated)
		if err_2 != nil {
			t.Error(mutated, err_2)
		}
	})
}