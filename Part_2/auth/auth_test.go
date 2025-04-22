package auth

import (
	"encoding/base64"
	"math/rand"
	"testing"
	"time"
)

func customMutator(input []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	switch rand.Intn(5) {
	case 0:
		return append(input, []byte(`","injected":"value"`)...)
	case 1:
		return append([]byte(`{"`), input...)
	case 2:
		return []byte(`"username":"admin","password":"'; DROP TABLE users; --"`)
	case 3:
		return append(input, input...)
	default:
		return input
	}
}

func FuzzParseSingleAuthHeader(f *testing.F) {
        f.Add([]byte(`"username"`))
		f.Add([]byte(`"password"`))
		f.Add([]byte(`{`))
		f.Add([]byte(`}`))
		f.Add([]byte(`,`))
		f.Add([]byte(`:`))
		f.Add([]byte(`""`))
        f.Fuzz(func(t *testing.T, input []byte) {
			mutated := customMutator(input)
			parseSingleAuthHeader(base64.URLEncoding.EncodeToString(mutated))
        })
}

func FuzzParseMultiAuthHeader(f *testing.F) {
	f.Add([]byte(`"username"`))
	f.Add([]byte(`"password"`))
	f.Add([]byte(`{`))
	f.Add([]byte(`}`))
	f.Add([]byte(`,`))
	f.Add([]byte(`:`))
	f.Add([]byte(`""`))
	f.Add([]byte(`https`))
	f.Add([]byte(`//`))
	f.Add([]byte(`.`))
	f.Fuzz(func(t *testing.T, input []byte) {
		mutated := customMutator(input)
		parseMultiAuthHeader(base64.URLEncoding.EncodeToString(mutated))
	})
}