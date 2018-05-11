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

func NewParser(fset *position.FileSet, errs *errors.Errors, filename string, src []rune) *parser {
	file := fset.AddFile(filename, fset.Base(), len(src))
	p := &parser{
		fset:    fset,
		scanner: scanner.NewScanner(file, src),
		errs:    errs,
	}
	p.scan()
	return p
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

		pe := s.parseDefine()
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
				//				s.errors(fmt.Errorf("Invalid expr '%v'", s.val))
				//				s.scanner.SkipError()
				s.scan()
			}
		}
	}
}
