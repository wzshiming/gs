package scanner

func (s *Scanner) scanString() string {
	off := s.off - 1
	ch := s.ch
	s.next()
	for ch != s.ch {
		s.next()
	}
	s.next()
	return string(s.buf[off : s.off-1])
}
