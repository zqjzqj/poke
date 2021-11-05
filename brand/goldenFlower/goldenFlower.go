package goldenFlower

import "github/zqjzqj/poker/brand"

type GoldenFlower struct {
	*brand.Brands
}

func NewGoldenFlower() *GoldenFlower {
	bs := brand.NewBrands(52)
	for _, decor := range brand.GetDecors() {
		for i := 1; i <= 13; i++ {
			b, _ := brand.NewBrand(uint8(i), decor)
			_ = bs.Add(b)
		}
	}
	return &GoldenFlower{
		Brands: bs,
	}
}
