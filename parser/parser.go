package parser

import (
	"fmt"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/scanner"
	"github.com/wzshiming/gs/token"
)

type parser struct {
	scanner *scanner.Scanner
	fset    *position.FileSet
	errs    *errors.Errors

	tok token.Token
	val string
	pos position.Pos
}

func NewParser(fset *position.FileSet, filename string, src []rune) *parser {
	file := fset.AddFile(filename, 1, len(src))
	p := &parser{
		fset:    fset,
		scanner: scanner.NewScanner(file, src),
		errs:    &errors.Errors{},
	}
	p.scan()
	return p
}

func (s *parser) Err() error {
	if s.errs.Len() == 0 {
		return nil
	}
	return s.errs
}

func (s *parser) scan() {
	var err error
	s.pos, s.tok, s.val, err = s.scanner.Scan()
	if err != nil {
		s.errors(err)
	}
}

func (s *parser) errors(err error) {
	s.errorsPos(s.pos, err)
}

func (s *parser) errorsPos(pos position.Pos, err error) {
	s.errs.Append(s.fset.Position(pos), err)
}

func (s *parser) Parse() []ast.Expr {
	ex := s.parse()
	if s.tok != token.EOF {
		s.errors(fmt.Errorf("Early exit '%v'", s.val))
	}
	return ex
}

func (s *parser) parse() []ast.Expr {
	ex := []ast.Expr{}
	for {

		pe := s.parseExpr()
		if pe != nil {
			ex = append(ex, pe)
		}

		switch s.tok {
		case token.EOF:
			return ex
		case token.RPAREN, token.RBRACK, token.RBRACE:
			return ex
		default:
			if pe == nil {
				s.errors(fmt.Errorf("Invalid expr '%v'", s.val))
				s.scanner.SkipError()
				s.scan()
			}
		}
	}
}

func (s *parser) parseExpr() ast.Expr {
	return s.parseBinaryExpr(1)
}

func (s *parser) parsePreUnaryExpr() (expr ast.Expr) {
	tok := s.tok
	pos := s.pos

	switch {
	case tok.IsOperator():
		switch s.tok {
		case token.SEMICOLON:
			s.scan()
			return s.parsePreUnaryExpr()
		case token.ADD, token.SUB, token.ELLIPSIS:
			s.scan()
			expr = &ast.OperatorPreUnary{
				Pos: pos,
				Op:  tok,
				X:   s.parsePreUnaryExpr(),
			}

		case token.RPAREN, token.RBRACE:
			// return nil
		case token.LPAREN:
			s.scan()
			b := s.parseExpr()
			if s.tok != token.RPAREN {
				s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
			}
			s.scan()
			expr = b
		case token.LBRACE:
			s.scan()
			b := s.parse()
			if s.tok != token.RBRACE {
				s.errorsPos(pos, fmt.Errorf("The parentheses are not closed '%s'", tok))
			}
			s.scan()

			expr = &ast.BraceExpr{
				Pos:  pos,
				List: b,
			}
		default:
			s.errors(fmt.Errorf("Undefined unary expr %v", s.val))
		}
	case tok.IsKeywork():
		switch tok {
		case token.IF:
			s.scan()
			cond := s.parseExpr()
			body := s.parseExpr()
			var els ast.Expr
			for s.tok == token.SEMICOLON {
				s.scan()
			}
			if s.tok == token.ELSE {
				s.scan()
				els = s.parseExpr()
			}
			expr = &ast.IfExpr{
				Pos:  pos,
				Cond: cond,
				Body: body,
				Else: els,
			}
		case token.FUNC:
			s.scan()
			name := s.parseExpr()
			args := s.parseExpr()
			body := s.parseExpr()
			expr = &ast.FuncExpr{
				Pos:  pos,
				Name: name,
				Args: args,
				Body: body,
			}
		case token.RETURN:
			s.scan()
			ret := s.parseExpr()
			expr = &ast.ReturnExpr{
				Pos: pos,
				Ret: ret,
			}
		default:
			s.errors(fmt.Errorf("Undefined keywork %v", s.tok))
		}

	default:
		switch tok {
		case token.EOF:
		case token.INVALID:
			s.errors(fmt.Errorf("Undefined value %v", s.val))
		default:
			b := &ast.Literal{
				Pos:   pos,
				Type:  s.tok,
				Value: s.val,
			}
			s.scan()
			expr = b
		}

	}

	if expr == nil {
		return
	}

loop:
	for {
		tok := s.tok
		pos := s.pos
		switch {
		case tok.IsOperator():
			switch s.tok {
			case token.INC, token.DEC, token.ELLIPSIS:

				expr = &ast.OperatorSufUnary{
					Pos: pos,
					Op:  tok,
					X:   expr,
				}
				s.scan()
			default:
				break loop
			}
		case tok.IsKeywork():
			break loop

		case tok.IsLiteral():
			expr = &ast.CallExpr{
				Pos:  pos,
				Name: expr,
				Args: s.parseExpr(),
			}
		default:

			break loop

		}
	}

	return expr
}

func (s *parser) parseBinaryExpr(pre int) ast.Expr {

	x := s.parsePreUnaryExpr()
	if x == nil {
		return x
	}

	for {
		op := s.tok
		pos := s.pos

		op2 := op.Precedence()
		if op2 < pre {
			break
		}
		s.scan()
		y := s.parseBinaryExpr(op2 + 1)

		switch op {
		case token.COMMA:
			if t, ok := x.(*ast.TupleExpr); ok {
				t.List = append(t.List, y)
			} else {
				x = &ast.TupleExpr{
					Pos:  pos,
					List: []ast.Expr{x, y},
				}
			}
		default:
			x = &ast.OperatorBinary{
				Pos: pos,
				X:   x,
				Op:  op,
				Y:   y,
			}
		}
	}
	return x
}
