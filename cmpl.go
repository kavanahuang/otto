package otto

import (
	"github.com/kavanahuang/otto/ast"
	"github.com/kavanahuang/otto/file"
)

type compiler struct {
	file    *file.File
	program *ast.Program
}
