package token

// looker is the symbolic looker
type looker struct {
	m   map[rune]*looker
	Tok Token
}

func newLooker() *looker {
	return &looker{}
}

func (l *looker) Add(r []rune, t Token) {
	if len(r) == 0 {
		l.Tok = t
		return
	}
	if l.m == nil {
		l.m = map[rune]*looker{}
	}
	if l.m[r[0]] == nil {
		l.m[r[0]] = newLooker()
	}
	l.m[r[0]].Add(r[1:], t)
	return
}

func (l *looker) GetRune(r rune) *looker {
	if l == nil {
		return l
	}
	return l.m[r]
}

func (l *looker) getRunes(r []rune) Token {
	if l == nil {
		return INVALID
	}
	if len(r) == 0 {
		return l.Tok
	}
	return l.m[r[0]].getRunes(r[1:])
}

func (l *looker) Get(s string) Token {
	r := []rune(s)
	return l.getRunes(r)
}
