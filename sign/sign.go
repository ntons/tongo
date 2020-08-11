package sign

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"net/url"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	fmt.Println("vim-go")
}

func s2b(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return
}

func sortValues(values url.Values) *bytes.Buffer {
	sortedKeys := make([]string, 0, len(values))
	for k := range values {
		if len(values.Get(k)) > 0 {
			sortedKeys = append(sortedKeys, k)
		}
	}
	sort.Strings(sortedKeys)
	buf := bytes.NewBuffer(nil)
	for i, k := range sortedKeys {
		if i > 0 {
			buf.WriteString("&")
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(values.Get(k))
	}
	return buf
}

func hashBytes(h hash.Hash, b []byte) string {
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MD5(values url.Values, key string) string {
	buf := sortValues(values)
	buf.WriteString(key)
	return hashBytes(md5.New(), buf.Bytes())
}

func SHA1(values url.Values, key string) string {
	buf := sortValues(values)
	buf.WriteString(key)
	return hashBytes(sha1.New(), buf.Bytes())
}

func SHA256(values url.Values, key string) string {
	buf := sortValues(values)
	buf.WriteString(key)
	return hashBytes(sha256.New(), buf.Bytes())
}

func HMACWithSHA1(values url.Values, key string) string {
	return hashBytes(hmac.New(sha1.New, s2b(key)), sortValues(values).Bytes())

}

func HMACWithSHA256(values url.Values, key string) string {
	return hashBytes(hmac.New(sha256.New, s2b(key)), sortValues(values).Bytes())
}
