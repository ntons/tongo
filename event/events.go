package event

type Events map[string]*Event

func (es Events) Watch(name string, callback CallbackFunc) CancelFunc {
	e, ok := es[name]
	if !ok {
		e = New()
		es[name] = e
	}
	return e.Watch(callback)
}

func (es Events) Emit(name string, x interface{}) {
	if e, ok := es[name]; ok {
		e.Emit(x)
	}
}
