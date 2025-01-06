package protobuf

import (
        "testing"
        "bytes"
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/apimachinery/pkg/runtime/schema"
)

func FuzzDecode(f *testing.F) {
        serializer := NewSerializer(nil, nil)
        into := &runtime.Unknown{}
        var gvk schema.GroupVersionKind
        f.Add([]byte{0x6b, 0x38, 0x73, 0x00})
        f.Fuzz(func(t *testing.T, originalData []byte) {
                prefixLen := len(serializer.prefix)
                switch {
                        case len(originalData) == 0:
                                t.Skip()
                        case len(originalData) < prefixLen || !bytes.Equal(serializer.prefix, originalData[:prefixLen]):
                                t.Skip()
                        case len(originalData) == prefixLen:
                                t.Skip()
                }
                _, _, err := serializer.Decode(originalData, &gvk, into)
                if err != nil && IsNotMarshalable(err) {
                        t.Error(err)
                }
        })
}
