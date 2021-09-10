package cli

type ReqAction struct {
	Send           Action
	Unmarshal      Action
	UnmarshalError Action
}

// Clear doc
func (ra *ReqAction) Clear() {
	ra.Unmarshal.Clear()
	ra.UnmarshalError.Clear()
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

// Action Name 大全
const (
	SendBasic             = "pfsense.SendBasicAction"
	UnmarshalAPIBasic     = "pfsense.unmarshalAPIBasic"
	UnmarshalBasic        = "pfsense.unmarshalBasic"
	UnmarshalEditCertPage = "pfsense.unmarshalEditCertPageResp"

	UnmarshalErrPageBasic    = "pfsense.unmarshalErrPageBasic"
	UnmarshalHaproxyList     = "pfsense.unmarshalHaproxyListResp"
	UnmarshalHaproxyNameList = "pfsense.unmarshalHaproxyNameListResp"
	UnmarshalIndex           = "pfsense.unmarshalIndexResp"
	UnmarshalLogin           = "pfsense.marshalLoginResp"
)

// Clear doc
func (a *Action) Clear() {
	a.list = a.list[0:0]
}

// Len doc
func (a *Action) Len() int {
	return len(a.list)
}

// PushBackNamed doc
func (a *Action) PushBackNamed(n NamedAction) {
	if cap(a.list) == 0 {
		a.list = make([]NamedAction, 0, 5)
	}
	a.list = append(a.list, n)
}

// PushFrontNamed doc
func (a *Action) PushFrontNamed(n NamedAction) {
	if cap(a.list) == len(a.list) {
		a.list = append([]NamedAction{n}, a.list...)
	} else {
		a.list = append(a.list, NamedAction{})
		copy(a.list[1:], a.list)
		a.list[0] = n
	}
}

// Remove doc
func (a *Action) Remove(n NamedAction) {
	a.RemoveByName(n.Name)
}

// RemoveByName doc
func (a *Action) RemoveByName(name string) {
	for i := 0; i < len(a.list); i++ {
		m := a.list[i]
		if m.Name == name {
			copy(a.list[i:], a.list[i+1:])
			a.list[len(a.list)-1] = NamedAction{}
			a.list = a.list[:len(a.list)-1]
			i--
		}
	}
}

// Run doc
func (a *Action) Run(r *Request) {
	for _, h := range a.list {
		h.Fn(r)
	}
}
