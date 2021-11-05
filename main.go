package main

import (
	"github/zqjzqj/poker/brand/goldenFlower"
	"github/zqjzqj/poker/logs"
)

func main() {
	gf := goldenFlower.NewGoldenFlower()
	gf.Shuffle()
	players := make([]*goldenFlower.Player, 0, 3)
	players = append(players, goldenFlower.NewPlayer(), goldenFlower.NewPlayer(), goldenFlower.NewPlayer())
	for _, p := range players {
		p.Brands[0] = gf.GetFirst()
		p.Brands[1] = gf.GetFirst()
		p.Brands[2] = gf.GetFirst()
	}

	for _, p := range players {
		logs.PrintlnSuccess(p.Brands)
	}
}
