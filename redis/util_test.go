package redis

import (
	"testing"
)

// 最大最小之间不能超过1
func checkDividedSlots(n int, a [][2]int) bool {
	// 必须有n个
	if len(a) != n {
		return false
	}
	// 必须以0开始，16383结束
	if a[0][0] != 0 || a[len(a)-1][1] != 16383 {
		return false
	}
	// 必须连续
	min, max := a[0][1]-a[0][0], a[0][1]-a[0][0]
	for i := 1; i < n; i++ {
		if a[i-1][1]+1 != a[i][0] {
			return false
		}
		if d := a[i][1] - a[i][0]; d < min {
			min = d
		} else if d > max {
			max = d
		}
	}
	// 每组数量只差不能大于1
	if max-min > 1 {
		return false
	}
	return true
}

func TestDivideSlots(t *testing.T) {
	for n := 1; n < 16384; n++ {
		a := divideSlots(n)
		if !checkDividedSlots(n, a) {
			t.Fatalf("divide error: %d,%v", n, a)
		}
	}
}
