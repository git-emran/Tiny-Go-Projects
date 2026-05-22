package main

import (
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	c := generativeart.NewCanva(600, 400)
	c.SetBackground(common.NavajoWhite)
	c.FillBackground()
	c.SetLineWidth(1.0)
	c.SetLineColor(common.Orange)
	c.Draw(arts.NewColorCircle2(30))
	c.ToPNG("circle.png")
}
