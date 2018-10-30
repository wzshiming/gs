package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"

	"github.com/wzshiming/gs/exec"
)

var profile = flag.String("profile", "", "profile file path")
var tree = flag.Bool("tree", false, "show ast tree")

func init() {
	flag.Parse()
}

func main() {

	if *profile != "" {
		f, _ := os.Create(*profile)
		defer f.Close()
		pprof.StartCPUProfile(f)     // start cpu profile，write to file
		defer pprof.StopCPUProfile() // stop profile
	}

	e := exec.NewExec()
	for _, filename := range flag.Args() {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		if *tree {
			v, err := e.Parse(filename, bytes.Runes(b))
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, v := range v {
				fmt.Println(v)
			}
		} else {
			v, err := e.Cmd(filename, bytes.Runes(b))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(v)
		}
	}
}
