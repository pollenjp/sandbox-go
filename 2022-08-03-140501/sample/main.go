package main

import (
	"fmt"
	"runtime"

	"github.com/pollenjp/sandbox-go/2022-08-03-140501/sample/path"
)

func main() {
	fmt.Println("Hello, world!")
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	fmt.Printf("Called from %s, line #%d, func: %v\n",
		file, line, runtime.FuncForPC(pc).Name())

	fpath := path.Path{Filepath: file}

	fmt.Println(fpath)
	fmt.Println(fpath.Base())
	fmt.Println(fpath.Stem())
	fmt.Println(fpath.Ext())

}
