package cli

type ReqAction struct {
	Send           Action
	Unmarshal      Action
	UnmarshalError Action
}

// Action doc
type Action struct {
	list []NamedAction
}

// NamedAction doc
type NamedAction struct {
	Name string
	Fn   func(*Request)
}

// PushBackNamed pushes named handler f to the back of the handler list.
func (l *Action) PushBackNamed(n NamedAction) {
	if cap(l.list) == 0 {
		l.list = make([]NamedAction, 0, 5)
	}
	l.list = append(l.list, n)
}

// PushFront pushes handler f to the front of the handler list.
func (l *Action) PushFront(f func(*Request)) {
	l.PushFrontNamed(NamedAction{"__anonymous", f})
}

// PushFrontNamed pushes named handler f to the front of the handler list.
func (l *Action) PushFrontNamed(n NamedAction) {
	if cap(l.list) == len(l.list) {
		// Allocating new list required
		l.list = append([]NamedAction{n}, l.list...)
	} else {
		// Enough room to prepend into list.
		l.list = append(l.list, NamedAction{})
		copy(l.list[1:], l.list)
		l.list[0] = n
	}
}

// Run executes all handlers in the list with a given request object.
func (l *Action) Run(r *Request) {
	for _, h := range l.list {
		h.Fn(r)
	}
}
