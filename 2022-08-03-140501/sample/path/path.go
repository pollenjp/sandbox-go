package path

import "path/filepath"

type Path struct {
	Filepath string
}

func (p Path) String() string {
	return p.Filepath
}

// parent directory
func (p Path) Abs() Path {
	pathStr, e := filepath.Abs(p.Filepath)
	if e != nil {
		panic(e)
	}
	return Path{Filepath: pathStr}
}

// filepath's basename
func (p Path) Base() string {
	return filepath.Base(p.Filepath)
}

// filename の拡張子を取得する
func (p Path) Ext() string {
	return filepath.Ext(p.Filepath)
}

// Join
func (p Path) Join(child string) Path {
	return Path{Filepath: filepath.Join(child)}
}

// filename exclude the extension
func (p Path) Parent() Path {
	return Path{Filepath: filepath.Dir(p.Filepath)}
}

// parent directory
func (p Path) Stem() string {
	var basename string = p.Base()
	return basename[:len(basename)-len(p.Ext())]
}
