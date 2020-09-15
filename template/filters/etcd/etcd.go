package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	//"github.com/etcd-io/etcd/client"
	"github.com/flosch/pongo2"
	"go.etcd.io/etcd/client"
)

func init() {
	pongo2.RegisterFilter("etcd", filterEtcdGet)
	pongo2.RegisterFilter("etcdget", filterEtcdGet)
	pongo2.RegisterFilter("etcdls", filterEtcdLs)
	pongo2.RegisterFilter("etcdjson", filterEtcdJSON)
}

func etcdGet(endpoints, key string) (_ *client.Response, err error) {
	cli, err := client.New(client.Config{
		Endpoints:               strings.Split(endpoints, ","),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})
	if err != nil {
		return
	}
	kapi := client.NewKeysAPI(cli)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return kapi.Get(ctx, key, nil)
}

// get value from etcd
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcdget:etcd_endpoints }}
func filterEtcdGet(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := etcdGet(param.String(), in.String())
	if err != nil {
		return nil, &pongo2.Error{OrigError: err}
	}
	if r.Node.Dir {
		return nil, &pongo2.Error{
			OrigError: fmt.Errorf("etcd node is directory"),
		}
	}
	return pongo2.AsSafeValue(r.Node.Value), nil
}

// list directory from etcd
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcdls:etcd_endpoints }}
func filterEtcdLs(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := etcdGet(param.String(), in.String())
	if err != nil {
		return nil, &pongo2.Error{OrigError: err}
	}
	if !r.Node.Dir {
		return nil, &pongo2.Error{
			OrigError: fmt.Errorf("etcd node is not directory"),
		}
	}
	var keys []string
	for _, node := range r.Node.Nodes {
		keys = append(keys, node.Key)
	}
	return pongo2.AsSafeValue(keys), nil
}

// get and unmarshal json structure from etcd
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcdjson:etcd_endpoints }}
func filterEtcdJSON(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := etcdGet(param.String(), in.String())
	if err != nil {
		return nil, &pongo2.Error{OrigError: err}
	}
	if r.Node.Dir {
		return nil, &pongo2.Error{
			OrigError: fmt.Errorf("etcd node is directory"),
		}
	}
	var x map[string]interface{}
	if err = json.Unmarshal([]byte(r.Node.Value), &x); err != nil {
		return nil, &pongo2.Error{OrigError: err}
	}
	return pongo2.AsSafeValue(x), nil
}
