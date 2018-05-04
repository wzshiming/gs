package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/token"
)

func (s *parser) parseKeywork() (expr ast.Expr) {
	tok := s.tok
	pos := s.pos

	switch tok {
	case token.FOR:
		fe := &ast.For{
			Pos: pos,
		}
		s.scan()
		if s.tok == token.SEMICOLON { // 开头是 ;
			s.scan()
			fe.Cond = s.parseExpr()

			if s.tok != token.SEMICOLON {
				s.errors(fmt.Errorf("No semicolon ends"))
				return nil
			}
			s.scan()
			if s.tok != token.LBRACE {
				fe.Next = s.parseExpr()
			}
		} else {
			initOrCond := s.parseExpr()

			if s.tok != token.SEMICOLON {
				fe.Cond = initOrCond
			} else {
				fe.Init = initOrCond
				s.scan()
				fe.Cond = s.parseExpr()
				if s.tok != token.SEMICOLON {
					s.errors(fmt.Errorf("No semicolon ends"))
					return nil
				}
				s.scan()
				if s.tok != token.LBRACE {
					fe.Next = s.parseExpr()
				}
			}
		}
		fe.Body = s.parseExpr()

		if s.tok == token.ELSE {
			s.scan()
			fe.Else = s.parseExpr()
		}
		return fe

	case token.IF:
		s.scan()
		init := s.parseExpr()
		cond := init

		if s.tok == token.SEMICOLON {
			s.scan()
			cond = s.parseExpr()
		} else {
			init = nil
		}

		body := s.parseExpr()
		var els ast.Expr
		if s.tok == token.ELSE {
			s.scan()
			els = s.parseExpr()
		}
		expr = &ast.If{
			Pos:  pos,
			Init: init,
			Cond: cond,
			Body: body,
			Else: els,
		}
		return expr
	case token.FUNC:
		s.scan()
		fun := s.parseExpr()
		body := s.parseExpr()
		expr = &ast.Func{
			Pos:  pos,
			Func: fun,
			Body: body,
		}
		return expr
	case token.RETURN:
		s.scan()
		ret := s.parseExpr()
		expr = &ast.Return{
			Pos: pos,
			Ret: ret,
		}
		return expr
	default:
		s.errors(fmt.Errorf("Undefined keywork %v", s.tok))
		return nil
	}
}
