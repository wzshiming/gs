package format

import (
	"bytes"

	"github.com/wzshiming/gs/printer"
)

func Format(src []rune) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	printer := printer.NewPrinter(buf)
	printer.Format(src)
	return buf.Bytes(), nil
}
