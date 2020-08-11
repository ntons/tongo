package event

type Events map[string]*Event

func (es Events) Watch(
	name string, callback CallbackFunc) (cancel CancelFunc) {
	e, ok := es[name]
	if !ok {
		e = New()
		es[name] = e
	}
	return e.Watch(callback)
}

func (es Events) Trigger(name string, x interface{}) {
	if e, ok := es[name]; ok {
		e.Trigger(x)
	}
}
