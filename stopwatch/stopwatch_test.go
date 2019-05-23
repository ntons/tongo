package stopwatch

import (
	"fmt"
	"testing"
	"time"
)

func TestElapsed(t *testing.T) {
	sw := New()
	time.Sleep(time.Second)
	fmt.Println(sw.Elapsed())
	time.Sleep(time.Second)
	fmt.Println(sw.Elapsed())
}
