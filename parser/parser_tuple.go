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
		es = append(es, s.parseBinary(1))
	}
	if len(es) == 1 {
		return es[0]
	}

	return &ast.Tuple{s.pos, es}
}
