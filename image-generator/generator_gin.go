package main

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

var DRAWINGS = map[string]generativeart.Engine{
	"maze":       arts.NewMaze(10),
	"julia":      arts.NewJulia(func(z complex128) complex128 { return z*z + complex(-0.1, 0.651) }, 40, 1.5, 1.5),
	"randcircle": arts.NewRandCicle(30, 80, 0.2, 2, 10, 30, true),
}
