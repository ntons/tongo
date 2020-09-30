package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/flosch/pongo2"
	"go.etcd.io/etcd/v3/client"
)

func init() {
	for _, e := range []struct {
		names  []string
		filter pongo2.FilterFunction
	}{
		{
			names: []string{
				"etcdget",
				"etcd_get",
				"etcd_v2_get",
			},
			filter: filterGet,
		},
		{
			names: []string{
				"etcdls",
				"etcd_ls",
				"etcd_v2_ls",
			},
			filter: filterLs,
		},
		{
			names: []string{
				"etcdjson",
				"etcd_json",
				"etcd_v2_json",
			},
			filter: filterJSON,
		},
	} {
		for _, name := range e.names {
			pongo2.RegisterFilter(name, e.filter)
		}
	}
}

func get(endpoints, key string) (_ *client.Response, err error) {
	cli, err := client.New(client.Config{
		Endpoints:               strings.Split(endpoints, ","),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 5 * time.Second,
	})
	if err != nil {
		return
	}
	kapi := client.NewKeysAPI(cli)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return kapi.Get(ctx, key, nil)
}

// get value
// names: etcdget, etcd_get, etcd_v2_get
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcdget:etcd_endpoints }}
func filterGet(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := get(param.String(), in.String())
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

// list directory
// names: etcdls, etcd_ls, etcd_v2_ls
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcdls:etcd_endpoints }}
func filterLs(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := get(param.String(), in.String())
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

// get and unmarshal json structure
// names: etcdjson, etcd_json, etcd_v2_json
// eg:
// {% set etcd_endpoints = 'http://127.0.0.1:2379,http://127.0.0.1:3379' %}
// {{ '/libra.net/librad/xxx'|etcd_json:etcd_endpoints }}
func filterJSON(
	in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	r, err := get(param.String(), in.String())
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
