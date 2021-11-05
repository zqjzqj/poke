package main

import (
	"github/zqjzqj/poker/brand/goldenFlower"
	"github/zqjzqj/poker/logs"
)

func main() {
	//初始化牌组
	gf := goldenFlower.NewGoldenFlower()

	//洗牌
	gf.Shuffle()

	//玩个玩家
	players := make([]*goldenFlower.Player, 0, 3)
	players = append(players, goldenFlower.NewPlayer(), goldenFlower.NewPlayer(), goldenFlower.NewPlayer())

	//发牌
	for _, p := range players {
		_ = p.AddBrand(gf.GetFirst())
		_ = p.AddBrand(gf.GetFirst())
		_ = p.AddBrand(gf.GetFirst())
	}

	//输出
	for _, p := range players {
		logs.PrintlnSuccess(p.GetBrands())
	}
}
