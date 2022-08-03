package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Path struct {
	Filepath string
}

// filepath's basename
func (p Path) Base() string {
	return filepath.Base(p.Filepath)
}

// filename の拡張子を取得する
func (p Path) Ext() string {
	return filepath.Ext(p.Filepath)
}

// filename exclude the extension
func (p Path) Stem() string {
	basename := p.Base()
	return basename[:len(basename)-len(p.Ext())]
}

func (p Path) String() string {
	return p.Filepath
}

func main() {
	fmt.Println("Hello, world!")
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	fmt.Printf("Called from %s, line #%d, func: %v\n",
		file, line, runtime.FuncForPC(pc).Name())

	fpath := Path{Filepath: file}

	fmt.Println(fpath)
	fmt.Println(fpath.Base())
	fmt.Println(fpath.Stem())
	fmt.Println(fpath.Ext())

}
