package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Options struct {
	// 集群节点配置，注意顺序
	NodeOptions []*redis.Options
}

func (o *Options) getAddrs() (a []string) {
	for _, node := range o.NodeOptions {
		a = append(a, node.Addr)
	}
	return
}

func (o *Options) getNewClient() func(*redis.Options) *redis.Client {
	return func(_o *redis.Options) *redis.Client {
		for _, node := range o.NodeOptions {
			if node.Addr == _o.Addr {
				return redis.NewClient(node)
			}
		}
		return nil
	}
}
func (o *Options) getClusterSlots() func(context.Context) ([]redis.ClusterSlot, error) {
	return func(context.Context) ([]redis.ClusterSlot, error) {
		n := len(o.NodeOptions)
		d := 16384 / n
		slots := make([]redis.ClusterSlot, 0, n)
		for i, node := range o.NodeOptions {
			slots = append(slots, redis.ClusterSlot{
				Start: d * i,
				End:   d*i + d - 1,
				Nodes: []redis.ClusterNode{{Addr: node.Addr}},
			})
		}
		slots[n-1].End = 16384
		return slots, nil
	}
}
