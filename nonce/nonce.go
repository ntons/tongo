package nonce

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const table = "0123456789abcdefghijklmnopqrstuvwxyz"

func Random(length int) string {
	max := big.NewInt(int64(len(table)))
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, max)
		sb.WriteByte(table[n.Int64()])
	}
	return sb.String()
}
