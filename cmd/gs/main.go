package main

import (
	"flag"
	"fmt"

	"github.com/wzshiming/gs/exec"
)

func main() {
	flag.Parse()
	e := exec.NewExec()
	for _, v := range flag.Args() {
		v, err := e.File(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
	}
}
