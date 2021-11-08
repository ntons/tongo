package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func xTestClient(t *testing.T, cli Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := cli.Ping(ctx).Err(); err != nil {
		t.Fatalf("failed to ping: %v", err)
	}
	if err := cli.Del(ctx,
		"0{hello}0",
		"1{hello}1",
		"0{world}0",
		"1{world}1",
	).Err(); err != nil {
		t.Fatalf("failed to delete: %v", err)
	}
	if err := cli.Set(ctx, "0{hello}0", "hello0", 0).Err(); err != nil {
		t.Fatalf("failed to set: %v", err)
	}
	if err := cli.Set(ctx, "1{hello}1", "hello1", 0).Err(); err != nil {
		t.Fatalf("failed to set: %v", err)
	}
	if err := cli.Set(ctx, "0{world}0", "world0", 0).Err(); err != nil {
		t.Fatalf("failed to set: %v", err)
	}
	if err := cli.Set(ctx, "1{world}1", "world1", 0).Err(); err != nil {
		t.Fatalf("failed to set: %v", err)
	}
}

func TestSingleClient(t *testing.T) {
	s := "redis://localhost:6379?dial_timeout=3&db=1&read_timeout=6s&max_retries=2"
	o, err := ParseURL(s)
	if err != nil {
		t.Fatalf("failed to parse url: %v", err)
	}
	cli := NewClient(o)
	if _, ok := cli.(*redis.Client); !ok {
		t.Fatalf("expect redis.Client but not")
	}
	xTestClient(t, cli)
}

func TestClusterClient(t *testing.T) {
	s := "redis://localhost:6379,localhost:6479,localhost:6579/2?dial_timeout=3&read_timeout=6s&max_retries=2"
	o, err := ParseURL(s)
	if err != nil {
		t.Fatalf("failed to parse url: %v", err)
	}
	cli := NewClient(o)
	if _, ok := cli.(*redis.ClusterClient); !ok {
		t.Fatalf("expect redis.ClusterClient but not")
	}
	xTestClient(t, cli)
}
