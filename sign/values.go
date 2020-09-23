package sign

import (
	"bytes"
	"sort"
)

type KV struct{ K, V string }

type Values []KV

type kn struct {
	k string
	n int
}

func (vals Values) buffer() *bytes.Buffer {
	idx := make([]kn, 0, len(vals))
	for i, kv := range vals {
		idx = append(idx, kn{kv.K, i})
	}
	sort.Slice(idx, func(i, j int) bool { return idx[i].k < idx[j].k })
	buf := bytes.NewBuffer(nil)
	for i, e := range idx {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(e.k)
		buf.WriteByte('=')
		buf.WriteString(vals[e.n].V)
	}
	return buf
}
