package parser

import (
	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseDefine() ast.Expr {
	pos := s.pos
	x := s.parseTuple()
	switch op := s.tok; op {
	case token.DEFINE, token.ASSIGN, token.ADD_ASSIGN, token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN, token.POW_ASSIGN, token.REM_ASSIGN, token.AND_ASSIGN, token.OR_ASSIGN, token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
		s.scan()
		y := s.parseTuple()
		return &ast.Binary{
			Pos: pos,
			X:   x,
			Op:  op,
			Y:   y,
		}
	default:
		return x
	}
}
