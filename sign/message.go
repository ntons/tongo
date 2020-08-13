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
func MessageToValues(msg proto.Message) (vals url.Values) {
	return messageToValues("", msg.ProtoReflect())
}

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

func messageToValues(pre string, msg reflect.Message) (vals url.Values) {
	vals = url.Values{}
	msg.Range(func(fd reflect.FieldDescriptor, v reflect.Value) bool {
		if fd.IsMap() {
			joinValues(vals, mapToValues(pre, fd, v))
		} else if fd.IsList() {
			joinValues(vals, listToValues(pre, fd, v))
		} else {
			fullPath := joinFullPath(pre, fd.Name())
			joinValues(vals, fieldToValues(fullPath, fd, v))
		}
		return true
	})
	return
}

func listToValues(pre string, fd reflect.FieldDescriptor, v reflect.Value) (vals url.Values) {
	if !fd.IsList() {
		panic("not list")
	}
	vals = url.Values{}
	for i := 0; i < v.List().Len(); i++ {
		fullPath := joinFullPath(pre, fd.Name(), i)
		joinValues(vals, fieldToValues(fullPath, fd, v.List().Get(i)))
	}
	return
}

func mapToValues(pre string, fd reflect.FieldDescriptor, v reflect.Value) (vals url.Values) {
	if !fd.IsMap() {
		panic("not map")
	}
	vals = url.Values{}
	v.Map().Range(func(k reflect.MapKey, v reflect.Value) bool {
		fullPath := joinFullPath(pre, fd.Name(), k)
		joinValues(vals, fieldToValues(fullPath, fd.MapValue(), v))
		return true
	})
	return
}

func fieldToValues(fullPath string, fd reflect.FieldDescriptor, v reflect.Value) (vals url.Values) {
	vals = url.Values{}
	switch fd.Kind() {
	case reflect.BoolKind:
		if !v.Bool() {
			vals.Set(fullPath, "0")
		} else {
			vals.Set(fullPath, "1")
		}
	case reflect.EnumKind:
		vals.Set(fullPath, fmt.Sprintf("%d", v.Enum()))
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
		vals.Set(fullPath, v.String())
	case reflect.BytesKind:
		vals.Set(fullPath, fmt.Sprintf("%x", v.Bytes()))
	case reflect.MessageKind:
		joinValues(vals, messageToValues(fullPath, v.Message()))
	case reflect.GroupKind:
		// Note that this feature is deprecated
		// and should not be used when creating new message types
	}
	return
}
