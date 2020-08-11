package sign

import (
	"fmt"
	"net/url"
	"strings"

	"google.golang.org/protobuf/proto"
	reflect "google.golang.org/protobuf/reflect/protoreflect"
)

// {a: { b: "a.b", c: { d: "a.c.d" } }, e: "e"}
// a.b=xxx&a.c.d=xxx&e=xxx
func MessageToValues(msg proto.Message) (values url.Values) {
	return msgToValues("", msg.ProtoReflect())
}

// join v2 INTO v1
func joinValues(dest, src url.Values) {
	for key, value := range src {
		dest[key] = value
	}
}

func joinFullPath(a ...interface{}) string {
	sb := strings.Builder{}
	for _, v := range a {
		if sb.Len() > 0 {
			sb.WriteByte('.')
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}

func msgToValues(prefix string, msg reflect.Message) (values url.Values) {
	values = url.Values{}
	msg.Range(func(fd reflect.FieldDescriptor, v reflect.Value) bool {
		if fd.IsMap() {
			joinValues(values, mapToValues(prefix, fd, v))
		} else if fd.IsList() {
			joinValues(values, listToValues(prefix, fd, v))
		} else {
			fullPath := joinFullPath(prefix, fd.Name())
			joinValues(values, toValues(fullPath, fd, v))
		}
		return true
	})
	return
}

func listToValues(prefix string, fd reflect.FieldDescriptor, v reflect.Value) (values url.Values) {
	if !fd.IsList() {
		panic("not list")
	}
	values = url.Values{}
	for i := 0; i < v.List().Len(); i++ {
		fullPath := joinFullPath(prefix, fd.Name(), i)
		joinValues(values, toValues(fullPath, fd, v.List().Get(i)))
	}
	return
}

func mapToValues(prefix string, fd reflect.FieldDescriptor, v reflect.Value) (values url.Values) {
	if !fd.IsMap() {
		panic("not map")
	}
	values = url.Values{}
	v.Map().Range(func(k reflect.MapKey, v reflect.Value) bool {
		fullPath := joinFullPath(prefix, fd.Name(), k)
		joinValues(values, toValues(fullPath, fd.MapValue(), v))
		return true
	})
	return
}

func toValues(fullPath string, fd reflect.FieldDescriptor, v reflect.Value) (values url.Values) {
	values = url.Values{}
	switch fd.Kind() {
	case reflect.BoolKind:
		if !v.Bool() {
			values.Set(fullPath, "0")
		} else {
			values.Set(fullPath, "1")
		}
	case reflect.EnumKind:
		values.Set(fullPath, fmt.Sprintf("%d", v.Enum()))
	case reflect.Int32Kind:
		fallthrough
	case reflect.Sint32Kind:
		fallthrough
	case reflect.Uint32Kind:
		fallthrough
	case reflect.Int64Kind:
		fallthrough
	case reflect.Sint64Kind:
		fallthrough
	case reflect.Uint64Kind:
		fallthrough
	case reflect.Sfixed32Kind:
		fallthrough
	case reflect.Fixed32Kind:
		fallthrough
	case reflect.FloatKind:
		fallthrough
	case reflect.Sfixed64Kind:
		fallthrough
	case reflect.Fixed64Kind:
		fallthrough
	case reflect.DoubleKind:
		fallthrough
	case reflect.StringKind:
		values.Set(fullPath, v.String())
	case reflect.BytesKind:
		values.Set(fullPath, fmt.Sprintf("%x", v.Bytes()))
	case reflect.MessageKind:
		joinValues(values, msgToValues(fullPath, v.Message()))
	case reflect.GroupKind:
		// Note that this feature is deprecated
		// and should not be used when creating new message types
	}
	return
}
