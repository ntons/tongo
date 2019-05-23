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

func (x *Stopwatch) Elapsed() time.Duration {
	return time.Since(x.t)
}

func (x *Stopwatch) Reset() {
	x.t = time.Now()
}
