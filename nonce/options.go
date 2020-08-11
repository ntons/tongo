package nonce

import (
	"time"
)

type options struct {
	timeout time.Duration
}

type Option interface {
	apply(o *options)
}

func defaultOptions() *options {
	return &options{
		timeout: 24 * time.Hour,
	}
}

type funcOption struct {
	f func(o *options)
}

func (fo funcOption) apply(o *options) {
	if fo.f != nil {
		fo.f(o)
	}
}

func WithTimeout(d time.Duration) Option {
	return funcOption{func(o *options) { o.timeout = d }}
}
