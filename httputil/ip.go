package httputil

import (
	"context"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

var privateNets []*net.IPNet

func init() {
	a := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
	for _, s := range a {
		_, privateNet, err := net.ParseCIDR(s)
		if err != nil {
			panic("bad private address cidr: " + s)
		}
		privateNets = append(privateNets, privateNet)
	}
}

func isPrivateIp(s string) bool {
	ip := net.ParseIP(s)
	if ip.IsLoopback() {
		return true
	}
	for _, privateNet := range privateNets {
		if privateNet.Contains(ip) {
			return true
		}
	}
	return false
}

// 获取HTTP请求的原始客户端IP
func GetRemoteIpFromContext(ctx context.Context) (ip string) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ip = getRemoteIp(md)
	}
	if ip == "" {
		if p, ok := peer.FromContext(ctx); ok {
			ip = strings.SplitN(p.Addr.String(), ":", 2)[0]
		}
	}
	return
}

func GetRemoteIpFromRequest(req *http.Request) (ip string) {
	if ip = getRemoteIp(req.Header); ip == "" {
		ip = strings.SplitN(req.RemoteAddr, ":", 2)[0]
	}
	return
}

func splitValues(s string) (v []string) {
	v = append(v, strings.Split(s, ",")...)
	for i, x := range v {
		v[i] = strings.TrimSpace(x)
	}
	return
}
func getValues(m map[string][]string, k string) ([]string, bool) {
	if v, ok := m[k]; ok && len(v) > 0 {
		return splitValues(v[0]), true
	}
	if v, ok := m[strings.ToLower(k)]; ok && len(v) > 0 {
		return splitValues(v[0]), true
	}
	return nil, false
}
func getRemoteIp(m map[string][]string) (ip string) {
	if v, ok := getValues(m, "X-Envoy-External-Address"); ok {
		// 获取经过envoy计算的外部ip
		// 这地方应该不会有多个值
		if len(v) > 0 {
			ip = v[0]
		}
	}
	if ip == "" {
		if v, ok := getValues(m, "X-Forwarded-For"); ok {
			// 经过某种代理转发的ip
			// 从右到左去掉内网IP，第一个公网IP当作地址
			for i := len(v) - 1; i >= 0; i-- {
				if !isPrivateIp(v[i]) {
					ip = v[i]
					break
				}
			}
		}
	}
	return
}
