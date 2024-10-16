package buy

import (
	"fmt"
)

func Buy() {
	var cookiesShop Shop = Shop{100, 5, 0}
	var queue []Buyer = []Buyer{{"Олег", 100}, {"Вова", 0}, {"акула", 1000}}
	var Buyers []Buyer = []Buyer{}

	queue[0], cookiesShop = queue[0].Buy(cookiesShop)

	Buyers = append(Buyers, queue[0])
	queue = queue[1:]

	fmt.Println("очередь:", queue)
	fmt.Println("покупатели:", Buyers)
	fmt.Println("магазин печенек:", cookiesShop)
}

type Shop struct {
	quantitySale int
	price        int
	money        int
}

type Buyer struct {
	name  string
	money int
}

func (b Buyer) Buy(s Shop) (Buyer, Shop) {
	if b.money > s.price {
		b.money -= s.price
		s.quantitySale -= 1
	}
	return b, s
}
