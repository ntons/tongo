package httputil

import (
	"testing"
)

func TestGetRemoteIp(t *testing.T) {
	expect := func(m map[string][]string, ip string) {
		if s := getRemoteIp(m); s != ip {
			t.Fatalf("expect \"%s\", got \"%s\", %v", ip, s, m)
		}
	}

	expect(map[string][]string{}, "")
	// 全内网
	expect(map[string][]string{
		"X-Forwarded-For": {"192.168.10.10,172.18.10.10"},
	}, "192.168.10.10")
	// 第一个外网
	expect(map[string][]string{
		"X-Forwarded-For": {"192.168.10.10,123.10.10.10,172.18.10.10"},
	}, "123.10.10.10")
	// 第一个
	expect(map[string][]string{
		"X-Envoy-Internal": {"true"},
		"X-Forwarded-For":  {"123.10.10.10,172.18.10.10"},
	}, "123.10.10.10")
	// external
	expect(map[string][]string{
		"X-Envoy-External-Address": {"123.10.10.10"},
	}, "123.10.10.10")
}
