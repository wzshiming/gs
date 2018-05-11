package scanner

func (s *Scanner) scanIdent() string {
	off := s.off - 1
loop:
	for {
		switch s.ch {
		case '+', '-', '*', '/', '%',
			'^', '&', '|',
			'\n', '\t', '\r', ' ', '\\',
			'(', ')', '[', ']', '{', '}',
			',', '.', ';', ':', -1:
			break loop
		}
		s.next()
	}
	return string(s.buf[off : s.off-1])
}
