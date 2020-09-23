package sign

import (
	"testing"
)

func TestProtoToValues(t *testing.T) {
	msg := &Foo{
		X1: []string{"x1.0", "x1.1"},
		X3: 1024,
		X4: map[string]string{
			"foo": "x4.foo",
			"bar": "x4.bar",
		},
		X5Xyz: "x5_xyz",
	}
	expected := "x1.0=x1.0&x1.1=x1.1&x3=1024&x4.bar=x4.bar&x4.foo=x4.foo&x5_xyz=x5_xyz"
	vals := ProtoToValues(msg)
	if s := vals.buffer().String(); s != expected {
		t.Fatalf("unexpected buffer result: %s", s)
	}
}
