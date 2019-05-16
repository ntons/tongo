package leaderboard

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -letl
#include "skiplist.h"
*/
import "C"

type Leaderboard struct {
}

func New() *Leaderboard {
	return &Leaderboard{}
}
