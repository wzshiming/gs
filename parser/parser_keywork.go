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
	case token.MAP:
		me := &ast.Map{
			Pos: pos,
		}
		s.scan()
		body := s.parseUnary()
		me.Body = body
		return me
	case token.FOR:
		fe := &ast.For{
			Pos: pos,
		}
		s.scan()
		if s.tok == token.SEMICOLON { // 开头是 ;
			s.scan()
			fe.Cond = s.parseDefine()

			if s.tok != token.SEMICOLON {
				s.errors(fmt.Errorf("No semicolon ends"))
				return nil
			}
			s.scan()
			if s.tok != token.LBRACE {
				fe.Next = s.parseDefine()
			}
		} else {
			initOrCond := s.parseDefine()

			if s.tok != token.SEMICOLON {
				fe.Cond = initOrCond
			} else {
				fe.Init = initOrCond
				s.scan()
				fe.Cond = s.parseDefine()
				if s.tok != token.SEMICOLON {
					s.errors(fmt.Errorf("No semicolon ends"))
					return nil
				}
				s.scan()
				if s.tok != token.LBRACE {
					fe.Next = s.parseDefine()
				}
			}
		}
		fe.Body = s.parseDefine()

		if s.tok == token.ELSE {
			s.scan()
			fe.Else = s.parseDefine()
		}
		return fe

	case token.IF:
		s.scan()
		init := s.parseDefine()
		cond := init

		if s.tok == token.SEMICOLON {
			s.scan()
			cond = s.parseDefine()
		} else {
			init = nil
		}

		body := s.parseDefine()
		var els ast.Expr
		if s.tok == token.ELSE {
			s.scan()
			els = s.parseDefine()
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
		fun := s.parseDefine()
		body := s.parseDefine()
		expr = &ast.Func{
			Pos:  pos,
			Func: fun,
			Body: body,
		}
		return expr
	case token.RETURN:
		s.scan()
		ret := s.parseDefine()
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
