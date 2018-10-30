package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/wzshiming/gs/format"
)

func init() {
	flag.Parse()
}

func main() {

	for _, filename := range flag.Args() {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		out, err := format.Format(bytes.Runes(b))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(out))
	}
}
