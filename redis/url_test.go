package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestParseURL(t *testing.T) {
	a := []string{"host1:6379", "host2:6479", "host3:6579"}
	s := fmt.Sprintf(
		"redis://user:password@%s,%s,%s/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2",
		a[0], a[1], a[2])
	o, err := ParseURL(s)
	if err != nil {
		t.Fatalf("failed to parse url: %v", err)
	}
	for i, o := range o.NodeOptions {
		if o.Username != "user" ||
			o.Password != "password" ||
			o.Addr != a[i] ||
			o.DB != 1 ||
			o.DialTimeout != 3*time.Second ||
			o.ReadTimeout != 6*time.Second ||
			o.MaxRetries != 2 {
			t.Fatalf("invalid parsed value: %q", o.Addr)
		}
	}
}
