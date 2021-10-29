package redis

import (
	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type Client = redis.Cmdable

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
