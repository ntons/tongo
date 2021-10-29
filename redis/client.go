package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type Cmd = redis.Cmd
type IntCmd = redis.IntCmd
type BoolCmd = redis.BoolCmd
type StatusCmd = redis.StatusCmd
type StringCmd = redis.StringCmd

type Client interface {
	Close() error
	Ping(ctx context.Context) *StatusCmd
	Get(context.Context, string) *StringCmd
	Set(context.Context, string, interface{}, time.Duration) *StatusCmd
	SetNX(context.Context, string, interface{}, time.Duration) *BoolCmd
	Del(context.Context, ...string) *IntCmd
	EvalSha(context.Context, string, []string, ...interface{}) *Cmd
	ScriptLoad(context.Context, string) *StringCmd
	ScriptFlush(context.Context) *StatusCmd
}

var _ Client = (*redis.Client)(nil)
var _ Client = (*redis.ClusterClient)(nil)

func NewClient(o *Options) Client {
	// 没节点直接抛异常
	if len(o.NodeOptions) == 0 {
		panic("require one node at lease")
	}
	// 单节点返回普通的Client
	if len(o.NodeOptions) == 1 {
		return redis.NewClient(o.NodeOptions[0])
	}
	// 多节点返回集群Client
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        o.getAddrs(),
		NewClient:    o.getNewClient(),
		ClusterSlots: o.getClusterSlots(),
		ReadOnly:     false,
	})
}
