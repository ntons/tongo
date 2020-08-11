// 本地（进程内）实现

package nonce

import (
	"container/list"
	"sync"
	"time"
)

type localNonceTime struct {
	nonce string
	time  time.Time
}

type LocalChecker struct {
	o *options
	// 保护以下容器
	mu sync.Mutex
	// 查找容器，处理匹配
	m map[string]struct{}
	// 顺序容器，处理超时
	l *list.List
}

func NewLocalChecker(opts ...Option) *LocalChecker {
	o := defaultOptions()
	for _, opt := range opts {
		opt.apply(o)
	}
	return &LocalChecker{
		o: o,
		m: make(map[string]struct{}),
		l: list.New(),
	}
}

func (c *LocalChecker) Check(nonce string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	// 移除超时数据
	t := time.Now().Add(-c.o.timeout)
	for {
		e := c.l.Front()
		if e == nil {
			break
		}
		v := e.Value.(*localNonceTime)
		if v.time.After(t) {
			break
		}
		delete(c.m, v.nonce)
		c.l.Remove(e)
	}
	// 检查是否已存在
	if _, ok := c.m[nonce]; ok {
		return false
	}
	c.m[nonce] = struct{}{}
	c.l.PushBack(&localNonceTime{nonce, time.Now()})
	return true
}
