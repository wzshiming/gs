package printer

import (
	"bytes"
	"fmt"
	"io"

	"github.com/wzshiming/gs/ast"
	"github.com/wzshiming/gs/errors"
	"github.com/wzshiming/gs/parser"
	"github.com/wzshiming/gs/position"
)

type Printer struct {
	output io.Writer
	level  int
	fset   *position.FileSet
	errs   *errors.Errors
}

func NewPrinter(output io.Writer) *Printer {
	return &Printer{
		output: output,
		fset:   position.NewFileSet(),
		errs:   errors.NewErrors(),
	}
}

func (p *Printer) Format(src []rune) ([]byte, error) {
	par := parser.NewParser(p.fset, p.errs, "", src)
	exprs := par.Parse()
	if p.errs.Len() != 0 {
		return nil, p.errs
	}
	buf := bytes.NewBuffer(nil)
	p.Print(exprs)
	return buf.Bytes(), nil
}

func (p *Printer) printf(format string) (int, error) {
	return io.WriteString(p.output, format)
}

func (p *Printer) changeTab(s int) {
	p.level += s
	if p.level < 0 {
		p.level = 0
	}
}

func (p *Printer) newLine() {
	p.output.Write([]byte{'\n'})
	for i := 0; i != p.level; i++ {
		p.output.Write([]byte{'\t'})
	}
}

func (p *Printer) Print(e []ast.Expr) {
	p.printNodes(e)
}

func (p *Printer) printNodes(e []ast.Expr) {
	for _, e := range e {
		p.newLine()
		p.printNode(e)
	}
}

func (p *Printer) printBrace(e ast.Expr) {
	if _, ok := e.(*ast.Brace); ok {
		p.printNode(e)
	} else {
		p.printf("{")
		p.changeTab(1)
		p.printNode(e)
		p.changeTab(-1)
		p.newLine()
		p.printf("}")
	}
}

func (p *Printer) printNode(e ast.Expr) {
	switch t := e.(type) {
	case *ast.Literal:
		p.printf(t.Value)
	case *ast.Binary:
		p.printNode(t.X)
		p.printf(" ")
		p.printf(t.Op.String())
		p.printf(" ")
		p.printNode(t.Y)
	case *ast.UnaryPre:
		p.printf(t.Op.String())
		p.printNode(t.X)
	case *ast.UnarySuf:
		p.printNode(t.X)
		p.printf(t.Op.String())
	case *ast.If:
		p.printf("if ")
		if t.Init != nil {
			p.printNode(t.Init)
			p.printf("; ")
		}
		p.printNode(t.Cond)
		p.printf(" ")
		p.printBrace(t.Body)
		if t.Else != nil {
			p.printf("else ")
			p.printBrace(t.Else)
		}
	case *ast.Brack:
		p.printNode(t.X)
		p.printf("[")
		p.printNode(t.Y)
		p.printf("]")
	case *ast.Labeled:
		// TODO:

	case *ast.Break:
		// TODO:

	case *ast.Continue:
		// TODO:

	case *ast.For:
		p.printf("for ")
		if t.Init != nil || t.Next != nil {
			p.printNode(t.Init)
			p.printf("; ")
			p.printNode(t.Cond)
			p.printf("; ")
			p.printNode(t.Next)
		}
		p.printf(" ")
		p.printBrace(t.Body)
		if t.Else != nil {
			p.printf("else ")
			p.printBrace(t.Else)
		}
	case *ast.Brace:
		p.printf("{")
		p.changeTab(1)
		p.printNodes(t.List)
		p.changeTab(-1)
		p.newLine()
		p.printf("}")
	case *ast.Call:
		p.printNode(t.Name)
		p.printf(" ")
		p.printNode(t.Args)
	case *ast.Func:
		p.printf("func ")
		p.printNode(t.Func)
		p.printf(" ")
		p.printNode(t.Body)
	case *ast.Return:
		p.printf("return ")
		p.printNode(t.Ret)
	case *ast.Tuple:
		p.printf("(")
		for i, v := range t.List {
			p.printNode(v)
			if len(t.List)-1 != i {
				p.printf(", ")
			}
		}
		p.printf(")")
	case *ast.Map:
		p.printf("map ")
		p.printNode(t.Body)
	}
	p.errorsPos(e.GetPos(), fmt.Errorf("Undefined keyword processing"))
}

func (p *Printer) errorsPos(pos position.Pos, err error) {
	p.errs.Append(p.fset.Position(pos), err)
}
