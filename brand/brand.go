package brand

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	DecorSpades   uint8 = 0
	DecorHearts   uint8 = 1
	DecorClubs    uint8 = 2
	DecorDiamonds uint8 = 3
)

func BrandDecorMap(decor uint8) string {
	if decor == DecorHearts {
		return "红心"
	}
	if decor == DecorSpades {
		return "黑桃"
	}
	if decor == DecorClubs {
		return "梅花"
	}
	if decor == DecorDiamonds {
		return "方块"
	}
	return ""
}

func BrandNumberString(num uint8) string {
	if num > 10 {
		switch num {
		case 11:
			return "J"
		case 12:
			return "Q"
		case 13:
			return "K"
		case 14:
			return "JokerPoker"
		case 15:
			return "JokerPoker"
		}
	} else {
		if num == 1 {
			return "A"
		}
		return fmt.Sprintf("%d", num)
	}
	return ""
}

func GetDecors() []uint8 {
	return []uint8{DecorSpades, DecorHearts, DecorClubs, DecorDiamonds}
}

type Brand struct {
	Number uint8
	Decor  uint8
}

func NewBrand(number, decor uint8) (*Brand, error) {
	if number < 0 || number > 15 {
		return nil, fmt.Errorf("invalid number")
	}
	if BrandDecorMap(decor) == "" {
		return nil, fmt.Errorf("invalid decor")
	}
	return &Brand{
		Number: number,
		Decor:  decor,
	}, nil
}

func (b *Brand) String() string {
	if b.Number == 14 {
		return "小王"
	}
	if b.Number == 15 {
		return "大王"
	}
	return fmt.Sprintf("%s%s", BrandDecorMap(b.Decor), BrandNumberString(b.Number))
}

//b > b1 return 1
//b < b1 return -1
//b == b1 return 0
func (b *Brand) Compare(b1 *Brand) int8 {
	if b.Number == b1.Number {
		if b.Decor == b1.Decor {
			return 0
		}
		if b.Decor < b1.Decor {
			return 1
		} else {
			return -1
		}
	}
	if b.Number > b1.Number {
		return 1
	}
	return -1
}

type Brands struct {
	bs     []*Brand
	used   []*Brand
	length uint
	mutex  sync.Mutex
}

func NewBrands(length uint) *Brands {
	if length == 0 {
		length = 54
	}
	return &Brands{
		length: length,
		bs:     make([]*Brand, 0, length),
		used:   make([]*Brand, 0, length),
	}
}

func (bs *Brands) Add(b *Brand) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	if bs.length <= uint(len(bs.bs)+len(bs.used)) {
		return fmt.Errorf("not add, length already enough %d %d", bs.length, (uint(len(bs.bs) + len(bs.used))))
	}
	bs.bs = append(bs.bs, b)
	return nil
}

func (bs *Brands) Shuffle() {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	if len(bs.used) > 0 {
		bs.bs = append(bs.bs, bs.used...)
		bs.used = bs.used[:]
	}
	bln := len(bs.bs)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < bln; i++ {
		rn := r.Intn(bln)
		bs.bs[i], bs.bs[rn] = bs.bs[rn], bs.bs[i]
	}
}

func (bs *Brands) GetFirst() *Brand {
	bln := len(bs.bs)
	if bln == 0 {
		return nil
	}
	r := bs.bs[0]
	copy(bs.bs, bs.bs[1:])
	bs.bs = bs.bs[:bln-1]
	bs.used = append(bs.used, r)
	return r
}

func (bs *Brands) GetBrands() []*Brand {
	return bs.bs
}

func (bs *Brands) GetUsdBrands() []*Brand {
	return bs.used
}
