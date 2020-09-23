package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"reflect"
	"unsafe"

	"google.golang.org/protobuf/proto"
)

func s2b(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return
}

func MD5(vals Values, key string) string {
	buf := vals.buffer()
	buf.WriteString(key)
	return fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
}
func SHA1(vals Values, key string) string {
	buf := vals.buffer()
	buf.WriteString(key)
	return fmt.Sprintf("%x", sha1.Sum(buf.Bytes()))
}
func SHA256(vals Values, key string) string {
	buf := vals.buffer()
	buf.WriteString(key)
	return fmt.Sprintf("%x", sha256.Sum256(buf.Bytes()))
}
func HMACWithSHA1(vals Values, key string) string {
	h := hmac.New(sha1.New, s2b(key))
	h.Write(vals.buffer().Bytes())
	return fmt.Sprintf("%x", h.Sum(nil))
}
func HMACWithSHA256(vals Values, key string) string {
	h := hmac.New(sha256.New, s2b(key))
	h.Write(vals.buffer().Bytes())
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ProtoMD5(msg proto.Message, key string) string {
	return MD5(ProtoToValues(msg), key)
}
func ProtoSHA1(msg proto.Message, key string) string {
	return SHA1(ProtoToValues(msg), key)
}
func ProtoSHA256(msg proto.Message, key string) string {
	return SHA256(ProtoToValues(msg), key)
}
func ProtoHMACWithSHA1(msg proto.Message, key string) string {
	return HMACWithSHA1(ProtoToValues(msg), key)
}
func ProtoHMACWithSHA256(msg proto.Message, key string) string {
	return HMACWithSHA1(ProtoToValues(msg), key)
}
