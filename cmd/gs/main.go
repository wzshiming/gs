package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/wzshiming/gs/exec"
)

var profile = flag.String("profile", "", "profile file path")

func init() {
	flag.Parse()
}

func main() {

	if *profile != "" {
		f, _ := os.Create(*profile)
		defer f.Close()
		pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
		defer pprof.StopCPUProfile() // 结束profile
	}

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
