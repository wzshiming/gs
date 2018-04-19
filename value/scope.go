package value

type Scope struct {
	parent *Scope
	scope  map[string]Value
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
	}
}

func (p *Scope) NewChildScope() *Scope {
	return NewScope(p)
}

func (p *Scope) Get(name string) (Value, bool) {
	if p == nil {
		return nil, false
	}
	if len(p.scope) == 0 {
		return p.parent.Get(name)
	}
	v, ok := p.scope[name]
	if !ok {
		return p.parent.Get(name)
	}
	return v, ok
}

func (p *Scope) Set(name string, val Value) {
	if p == nil {
		return
	}
	if len(p.scope) == 0 {
		p.scope = map[string]Value{}
	}
	p.scope[name] = val
	return
}
