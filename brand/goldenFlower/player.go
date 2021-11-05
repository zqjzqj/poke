package goldenFlower

import (
	"fmt"
	"github/zqjzqj/poker/brand"
)

type Player struct {
	brands  [3]*brand.Brand
	bLength uint8
}

func NewPlayer() *Player {
	return &Player{brands: [3]*brand.Brand{}}
}

func (p *Player) AddBrand(b *brand.Brand) error {
	if p.bLength == 3 {
		return fmt.Errorf("not add, length already enough 3")
	}
	for k, v := range p.brands {
		if v != nil {
			if v.Compare(b) > 0 {
				p.brands[k], b = b, p.brands[k]
			}
		} else {
			p.brands[k] = b
			break
		}
	}
	p.bLength++
	return nil
}

func (p *Player) ClearBrands() {
	p.bLength = 0
	p.brands[0] = nil
	p.brands[1] = nil
	p.brands[2] = nil
}

func (p Player) GetBrands() [3]*brand.Brand {
	return p.brands
}

func (p Player) GetBrandsByPos(pos int) *brand.Brand {
	if pos > 3 || pos < 1 {
		return nil
	}
	return p.brands[pos-1]
}
