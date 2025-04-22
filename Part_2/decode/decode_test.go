package protobuf

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func mutateBytes(input []byte, prefix []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	mutations := []func([]byte) []byte{
		func(data []byte) []byte {
			return append(prefix, data...)
		},
		func(data []byte) []byte {
			return append(bytes.Repeat(prefix, 2), data...)
		},
		func(data []byte) []byte {
			suffix := make([]byte, rand.Intn(10)+1)
			rand.Read(suffix)
			return append(prefix, suffix...)
		},
		func(data []byte) []byte {
			return append(prefix, data[:rand.Intn(len(data))]...)
		},
		func(data []byte) []byte {
			return append(prefix, append([]byte{0, 0, 0}, data...)...)
		},
		func(data []byte) []byte {
			return prefix
		},
	}
	return mutations[rand.Intn(len(mutations))](input)
}

func FuzzDecode(f *testing.F) {
	serializer := NewSerializer(nil, nil)
	into := &runtime.Unknown{}
	var gvk schema.GroupVersionKind
	f.Add([]byte{0x6b, 0x38, 0x73, 0x00})
	f.Fuzz(func(t *testing.T, originalData []byte) {
		if len(originalData) == 0 {
			t.Skip()
		}
		mutated := mutateBytes(originalData, serializer.prefix)
		prefixLen := len(serializer.prefix)
		if len(mutated) < prefixLen || !bytes.Equal(serializer.prefix, mutated[:prefixLen]) {
			t.Skip()
		}
		if len(mutated) == prefixLen {
			t.Skip()
		}
		_, _, err := serializer.Decode(mutated, &gvk, into)
		if err != nil && IsNotMarshalable(err) {
			t.Error(err)
		}
	})
}