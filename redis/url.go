package redis

import (
	"net/url"
	"strings"

	"github.com/go-redis/redis/v8"
)

// Example: redis://user:password@host1:6379,host2:6479,host3:6579/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2
func ParseURL(redisURL string) (*Options, error) {
	u, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}
	o, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	r := &Options{}
	for _, s := range strings.Split(u.Host, ",") {
		clone := *o
		clone.Addr = s
		r.NodeOptions = append(r.NodeOptions, &clone)
	}
	return r, nil
}
