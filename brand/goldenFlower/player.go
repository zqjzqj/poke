package goldenFlower

import "github/zqjzqj/poker/brand"

type Player struct {
	Brands [3]*brand.Brand
}

func NewPlayer() *Player {
	return &Player{Brands: [3]*brand.Brand{}}
}

func (p *Player) AddBrands(b *brand.Brand) error {
	if p.Brands[0] == nil {
		p.Brands[0] = b
	}
	if p.Brands[0].Compare(b) < 0 {
		p.Brands[0], b = b, p.Brands[0]
	}
	if p.Brands[1].Compare(b) < 0 {
		p.Brands[1], b = b, p.Brands[1]
	}
	if p.Brands[2].Compare(b) < 0 {
		p.Brands[2], b = b, p.Brands[2]
	}
	return nil
}
