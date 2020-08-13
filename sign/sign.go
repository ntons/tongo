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
	"sort"

	"github.com/ntons/tongo/nocopy"
	"google.golang.org/protobuf/proto"
)

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

func hashWith(h hash.Hash, b []byte) string {
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func toValues(v interface{}) (vals url.Values) {
	if x, ok := v.(url.Values); ok {
		vals = x
	} else if x, ok := v.(proto.Message); ok {
		vals = MessageToValues(x)
	} else {
		panic(fmt.Errorf("invalid argument"))
	}
	return
}

func MD5(vals url.Values, key string) string {
	buf := sortValues(vals)
	buf.WriteString(key)
	return hashWith(md5.New(), buf.Bytes())
}
func SHA1(vals url.Values, key string) string {
	buf := sortValues(vals)
	buf.WriteString(key)
	return hashWith(sha1.New(), buf.Bytes())
}
func SHA256(vals url.Values, key string) string {
	buf := sortValues(vals)
	buf.WriteString(key)
	return hashWith(sha256.New(), buf.Bytes())
}
func HMACWithSHA1(vals url.Values, key string) string {
	return hashWith(
		hmac.New(sha1.New, nocopy.StringToBytes(key)),
		sortValues(vals).Bytes())
}
func HMACWithSHA256(vals url.Values, key string) string {
	return hashWith(
		hmac.New(sha256.New, nocopy.StringToBytes(key)),
		sortValues(vals).Bytes())
}

func MessageMD5(msg proto.Message, key string) string {
	return MD5(MessageToValues(msg), key)
}
func MessageSHA1(msg proto.Message, key string) string {
	return SHA1(MessageToValues(msg), key)
}
func MessageSHA256(msg proto.Message, key string) string {
	return SHA256(MessageToValues(msg), key)
}
func MessageHMACWithSHA1(msg proto.Message, key string) string {
	return HMACWithSHA1(MessageToValues(msg), key)
}
func MessageHMACWithSHA256(msg proto.Message, key string) string {
	return HMACWithSHA1(MessageToValues(msg), key)
}
