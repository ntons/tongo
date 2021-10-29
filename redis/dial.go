package redis

import "context"

func Dial(ctx context.Context, url string, dialOptArray ...DialOption) (_ Client, err error) {
	opts, err := ParseURL(url)
	if err != nil {
		return
	}
	cli := NewClient(opts)

	dialOpts := applyDialOptions(dialOptArray...)
	if dialOpts.WithPingTest {
		if err = cli.Ping(ctx).Err(); err != nil {
			return
		}
	}

	return cli, nil
}

type dialOptions struct {
	WithPingTest bool
}

func applyDialOptions(dialOpts ...DialOption) *dialOptions {
	var o dialOptions
	for _, _o := range dialOpts {
		if _o != nil {
			_o(&o)
		}
	}
	return &o
}

type DialOption func(*dialOptions)

// 执行Ping测试
func WithPingTest() DialOption {
	return func(o *dialOptions) {
		o.WithPingTest = true
	}
}
