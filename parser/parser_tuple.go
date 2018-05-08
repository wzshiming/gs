package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseTuple() ast.Expr {
	es := []ast.Expr{s.parseBinary(1)}

	for {
		if s.tok != token.COMMA {
			break
		}
		s.scan()
		b := s.parseBinary(1)
		if b == nil {
			break
		}
		es = append(es, b)
	}
	if len(es) == 1 {
		return es[0]
	}

	return &ast.Tuple{
		Pos:  s.pos,
		List: es,
	}
}
