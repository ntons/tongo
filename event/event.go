package event

import (
	"sort"
)

type CallbackFunc func(interface{})
type CancelFunc func()

type entry struct {
	id       int
	callback CallbackFunc
}

type Event struct {
	entries []*entry
}

func New() *Event {
	return &Event{}
}

func (e *Event) Watch(callback CallbackFunc) CancelFunc {
	id := 0
	if n := len(e.entries); n > 0 {
		id = e.entries[n-1].id + 1
	}
	e.entries = append(e.entries, &entry{
		id:       id,
		callback: callback,
	})
	return func() {
		i := sort.Search(len(e.entries), func(i int) bool {
			return e.entries[i].id >= id
		})
		if i < len(e.entries) && e.entries[i].id == id {
			e.entries = append(e.entries[:i], e.entries[i+1:]...)
		}
	}
}

func (e *Event) Emit(x interface{}) {
	for _, _e := range e.entries {
		if _e.callback != nil {
			_e.callback(x)
		}
	}
}
