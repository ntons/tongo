package stopwatch

import (
	"time"
)

type Stopwatch struct {
	t time.Time
}

func New() *Stopwatch {
	return &Stopwatch{t: time.Now()}
}

func (sw *Stopwatch) Elapsed() time.Duration {
	return time.Since(sw.t)
}

func (sw *Stopwatch) Reset() {
	sw.t = time.Now()
}
