package scanner

func (s *Scanner) scanNumber() string {
	off := s.off - 1
	for s.ch >= '0' && s.ch <= '9' {
		s.next()
	}
	if s.ch == '.' {
		s.next()
		for s.ch >= '0' && s.ch <= '9' {
			s.next()
		}
	}
	return string(s.buf[off : s.off-1])
}
