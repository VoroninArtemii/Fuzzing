package auth

import (
        "encoding/base64"
        "testing"
)

func FuzzParseSingleAuthHeader(f *testing.F) {
        f.Add([]byte(`{"username":"","password":""}`))
		f.Add([]byte(`{}`))
		f.Add([]byte(`,`))
		f.Add([]byte(`:`))
		f.Add([]byte(`""`))
		f.Add([]byte(`{"username":"frfjelorf","password":"ferjfherkivbh"}`))
		f.Add([]byte(`{"username":"freihfjeriof","password":"vfjbvjhdfbv"}`))
		f.Add([]byte(`{"username":"u1","password":"p1"}`))
        f.Fuzz(func(t *testing.T, input []byte) {
            parseSingleAuthHeader(base64.URLEncoding.EncodeToString(input))
        })
}

func FuzzParseMultiAuthHeader(f *testing.F) {
	f.Add([]byte(`{"username":"","password":""}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`,`))
	f.Add([]byte(`:`))
	f.Add([]byte(`""`))
	f.Add([]byte(`{"username":"frfjelorf","password":"ferjfherkivbh"}`))
	f.Add([]byte(`{"username":"u1","password":"p1"}`))
	f.Add([]byte(`{"https://index.docker.io/v1/":{"username":"u1","password":"p1"},` + `"quay.io/libpod":{"username":"u2","password":"p2"}}`))
	f.Fuzz(func(t *testing.T, input []byte) {
		parseMultiAuthHeader(base64.URLEncoding.EncodeToString(input))
	})
}
