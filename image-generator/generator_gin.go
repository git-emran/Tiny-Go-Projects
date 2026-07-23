package main

import (
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

var DRAWINGS = map[string]generativeart.Engine{
	"maze":       arts.NewMaze(10),
	"julia":      arts.NewJulia(func(z complex128) complex128 { return z*z + complex(-0.1, 0.651) }, 40, 1.5, 1.5),
	"randcircle": arts.NewRandCicle(30, 80, 0.2, 2, 10, 30, true),
	"blackhole":  arts.NewBlackHole(200, 400, 0.01),
	"janus":      arts.NewJanus(5, 10),
	"random":     arts.NewRandomShape(150),
	"silksky":    arts.NewSilkSky(15, 5),
	"circles":    arts.NewColorCircle(30),
}

func main() {
	drawMany(DRAWINGS)
}

func drawMany(drawings map[string]generativeart.Engine) {
	for k := range drawings {
		drawOne(k)
	}
}

func drawOne(art string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
}
