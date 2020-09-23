package sign

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// {a: { b: "a.b", c: { d: "a.c.d" } }, e: "e"}
// a.b=xxx&a.c.d=xxx&e=xxx
func ProtoToValues(msg proto.Message) (vals Values) {
	return messageToValues("", msg.ProtoReflect())
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

func messageToValues(pre string, msg protoreflect.Message) (vals Values) {
	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.IsMap() {
			vals = append(vals, mapToValues(pre, fd, v)...)
		} else if fd.IsList() {
			vals = append(vals, listToValues(pre, fd, v)...)
		} else {
			fullPath := joinFullPath(pre, fd.Name())
			vals = append(vals, fieldToValues(fullPath, fd, v)...)
		}
		return true
	})
	return
}

func listToValues(
	pre string, fd protoreflect.FieldDescriptor, v protoreflect.Value) (vals Values) {
	if !fd.IsList() {
		panic("not list")
	}
	for i := 0; i < v.List().Len(); i++ {
		fullPath := joinFullPath(pre, fd.Name(), i)
		vals = append(vals, fieldToValues(fullPath, fd, v.List().Get(i))...)
	}
	return
}

func mapToValues(
	pre string, fd protoreflect.FieldDescriptor, v protoreflect.Value) (vals Values) {
	if !fd.IsMap() {
		panic("not map")
	}
	v.Map().Range(func(k protoreflect.MapKey, v protoreflect.Value) bool {
		fullPath := joinFullPath(pre, fd.Name(), k)
		vals = append(vals, fieldToValues(fullPath, fd.MapValue(), v)...)
		return true
	})
	return
}

func fieldToValues(
	fullPath string, fd protoreflect.FieldDescriptor, v protoreflect.Value) (vals Values) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		if !v.Bool() {
			vals = append(vals, KV{fullPath, "0"})
		} else {
			vals = append(vals, KV{fullPath, "1"})
		}
	case protoreflect.EnumKind:
		vals = append(vals, KV{fullPath, fmt.Sprintf("%d", v.Enum())})
	case protoreflect.Int32Kind:
		fallthrough
	case protoreflect.Sint32Kind:
		fallthrough
	case protoreflect.Uint32Kind:
		fallthrough
	case protoreflect.Int64Kind:
		fallthrough
	case protoreflect.Sint64Kind:
		fallthrough
	case protoreflect.Uint64Kind:
		fallthrough
	case protoreflect.Sfixed32Kind:
		fallthrough
	case protoreflect.Fixed32Kind:
		fallthrough
	case protoreflect.FloatKind:
		fallthrough
	case protoreflect.Sfixed64Kind:
		fallthrough
	case protoreflect.Fixed64Kind:
		fallthrough
	case protoreflect.DoubleKind:
		fallthrough
	case protoreflect.StringKind:
		vals = append(vals, KV{fullPath, v.String()})
	case protoreflect.BytesKind:
		vals = append(vals, KV{fullPath, fmt.Sprintf("%x", v.Bytes())})
	case protoreflect.MessageKind:
		vals = append(vals, messageToValues(fullPath, v.Message())...)
	case protoreflect.GroupKind:
		// Note that this feature is deprecated
		// and should not be used when creating new message types
	}
	return
}

// via protojson
func ProtoToValues2(msg proto.Message) (vals Values) {
	b, err := protojson.Marshal(msg)
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}
	// traverse map as a tree
	for k, v := range m {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			vals = append(vals, traverseMap("", reflect.ValueOf(v))...)
		case reflect.Slice:
			vals = append(vals, traverseSlice("", reflect.ValueOf(v))...)
		default:
			vals = append(vals, KV{k, fmt.Sprintf("%v", v)})
		}
	}
	return
}

func traverseMap(prefix string, v reflect.Value) (vals Values) {
	iter := v.MapRange()
	for iter.Next() {
		fullPath := joinFullPath(prefix, iter.Key().String())
		vv := iter.Value()
		switch vv.Type().Kind() {
		case reflect.Map:
			vals = append(vals, traverseMap(fullPath, vv)...)
		case reflect.Slice:
			vals = append(vals, traverseSlice(fullPath, vv)...)
		default:
			vals = append(vals, KV{fullPath, fmt.Sprintf("%s", vv.Interface())})
		}
	}
	return
}

func traverseSlice(prefix string, v reflect.Value) (vals Values) {
	for i := 0; i < v.Len(); i++ {
		fullPath := joinFullPath(prefix, i)
		vv := v.Index(i)
		switch vv.Type().Kind() {
		case reflect.Map:
			vals = append(vals, traverseMap(fullPath, vv)...)
		default:
			vals = append(vals, KV{fullPath, fmt.Sprintf("%s", vv.Interface())})
		}
	}
	return
}
