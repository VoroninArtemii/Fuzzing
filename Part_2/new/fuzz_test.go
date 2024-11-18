package new

import (
        "testing"
        "k8s.io/apimachinery/pkg/runtime"
        "k8s.io/apimachinery/pkg/runtime/schema"
        "k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
)

type SimpleObject struct {
        Field1 string
        Field2 int32
}

func (o *SimpleObject) GetObjectKind() schema.ObjectKind {
        return &runtime.Unknown{}
}

func (o *SimpleObject) DeepCopyObject() runtime.Object {
        return &SimpleObject{
                Field1: o.Field1,
                Field2: o.Field2,
        }
}

func FuzzDecode(f *testing.F) {
        serializer := protobuf.NewSerializer(nil, nil)
        into := &SimpleObject{}
        f.Fuzz(func(t *testing.T, originalData []byte) {
                var gvk schema.GroupVersionKind
                _, _, err := serializer.Decode(originalData, &gvk, into)
                if err != nil && !protobuf.IsNotMarshalable(err) {
                        t.Errorf("Decode failed with error: %v", err)
                }
        })
}