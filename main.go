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
		_ = p.AddBrands(gf.GetFirst())
		_ = p.AddBrands(gf.GetFirst())
		_ = p.AddBrands(gf.GetFirst())
	}

	for _, p := range players {
		logs.PrintlnSuccess(p.GetBrands())
	}
}
