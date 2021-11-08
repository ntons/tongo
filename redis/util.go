package redis

// 把[0,16383]一共16384个数字平局分成n份
func divideSlots(n int) [][2]int {
	const N = 16384
	d := N / n
	m := N % n
	a := make([][2]int, n)
	for i, k := 0, 0; i < n; i++ {
		a[i][0] = k
		if i < m {
			k += d + 1
		} else {
			k += d
		}
		a[i][1] = k - 1
	}
	return a
}
