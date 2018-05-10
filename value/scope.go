package value

type Scope struct {
	parent *Scope
	scope  map[Value]Value
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
	}
}

func (p *Scope) NewChildScope() *Scope {
	return NewScope(p)
}

func (p *Scope) Get(name Value) (Value, bool) {
	v, ok := p.getScope(name)
	if ok {
		return v.scope[name], ok
	}
	return Nil, ok
}

func (p *Scope) getScope(name Value) (*Scope, bool) {
	if p == nil {
		return nil, false
	}
	if len(p.scope) == 0 {
		return p.parent.getScope(name)
	}
	_, ok := p.scope[name]
	if !ok {
		return p.parent.getScope(name)
	}
	return p, ok
}

func (p *Scope) Set(name Value, val Value) {
	v, ok := p.getScope(name)
	if ok {
		v.scope[name] = val
		return
	}
	p.SetLocal(name, val)
	return
}

func (p *Scope) SetLocal(name Value, val Value) {
	if p == nil {
		return
	}
	if len(p.scope) == 0 {
		p.scope = map[Value]Value{}
	}
	p.scope[name] = val
	return
}
