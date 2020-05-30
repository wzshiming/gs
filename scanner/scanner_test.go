package scanner

import (
	"testing"

	"github.com/wzshiming/gs/position"
	"github.com/wzshiming/gs/token"
)

func TestScanner_Scan(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		wantPos position.Pos
		wantTok token.Token
		wantVal string
		wantErr bool
	}{
		{
			args: args{
				code: "",
			},
			wantPos: 1,
			wantTok: token.EOF,
			wantVal: "",
		},
		{
			args: args{
				code: "1",
			},
			wantPos: 1,
			wantTok: token.NUMBER,
			wantVal: "1",
		},
		{
			args: args{
				code: "1.",
			},
			wantPos: 1,
			wantTok: token.NUMBER,
			wantVal: "1.",
		},
		{
			args: args{
				code: ".0",
			},
			wantPos: 1,
			wantTok: token.NUMBER,
			wantVal: ".0",
		},
		{
			args: args{
				code: ".01",
			},
			wantPos: 1,
			wantTok: token.NUMBER,
			wantVal: ".01",
		},
		{
			args: args{
				code: "0.01",
			},
			wantPos: 1,
			wantTok: token.NUMBER,
			wantVal: "0.01",
		},
		{
			args: args{
				code: "+",
			},
			wantPos: 1,
			wantTok: token.ADD,
			wantVal: "+",
		},
		{
			args: args{
				code: ".",
			},
			wantPos: 1,
			wantTok: token.PERIOD,
			wantVal: ".",
		},
		{
			args: args{
				code: "hello",
			},
			wantPos: 1,
			wantTok: token.IDENT,
			wantVal: "hello",
		},
		{
			args: args{
				code: "true",
			},
			wantPos: 1,
			wantTok: token.BOOL,
			wantVal: "true",
		},
		{
			args: args{
				code: "false",
			},
			wantPos: 1,
			wantTok: token.BOOL,
			wantVal: "false",
		},
		{
			args: args{
				code: "nil",
			},
			wantPos: 1,
			wantTok: token.NIL,
			wantVal: "nil",
		},
		{
			args: args{
				code: "Hello",
			},
			wantPos: 1,
			wantTok: token.IDENT,
			wantVal: "Hello",
		},
		{
			args: args{
				code: "'hello'",
			},
			wantPos: 1,
			wantTok: token.STRING,
			wantVal: "'hello'",
		},
		{
			args: args{
				code: "// hello",
			},
			wantPos: 9,
			wantTok: token.EOF,
			wantVal: "",
		},
		{
			args: args{
				code: "# hello",
			},
			wantPos: 8,
			wantTok: token.EOF,
			wantVal: "",
		},
		{
			args: args{
				code: "# \n10",
			},
			wantPos: 3,
			wantTok: token.SEMICOLON,
			wantVal: "\n",
		},
		{
			args: args{
				code: " 10",
			},
			wantPos: 2,
			wantTok: token.NUMBER,
			wantVal: "10",
		},
		{
			args: args{
				code: "好",
			},
			wantPos: 1,
			wantTok: token.IDENT,
			wantVal: "好",
		},
		{
			args: args{
				code: "\\",
			},
			wantPos: 1,
			wantTok: token.INVALID,
			wantVal: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := position.NewFileSet()
			code := []rune(tt.args.code)
			s := NewScanner(fs.AddFile(tt.name, fs.Base(), len(code)), code)
			gotPos, gotTok, gotVal, err := s.Scan()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPos != tt.wantPos {
				t.Errorf("Scan() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
			if gotTok != tt.wantTok {
				t.Errorf("Scan() gotTok = %v, want %v", gotTok, tt.wantTok)
			}
			if gotVal != tt.wantVal {
				t.Errorf("Scan() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}
