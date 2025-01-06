package auth

import (
        "encoding/base64"
        "testing"
)

func FuzzParseSingleAuthHeader(f *testing.F) {
        f.Add([]byte(`"username"`))
		f.Add([]byte(`"password"`))
		f.Add([]byte(`{`))
		f.Add([]byte(`}`))
		f.Add([]byte(`,`))
		f.Add([]byte(`:`))
		f.Add([]byte(`""`))
        f.Fuzz(func(t *testing.T, input []byte) {
            parseSingleAuthHeader(base64.URLEncoding.EncodeToString(input))
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
		parseMultiAuthHeader(base64.URLEncoding.EncodeToString(input))
	})
}
