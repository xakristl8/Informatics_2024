package structshop

func BuySell(b Buyer, s Shop) {
	b.Buy(s)
	s.Sell(b)
}

type Shop struct {
	quantitySale int
	price        int
	money        int
}

type Buyer struct {
	name    string
	money   int
	haveBuy int
}

func (b Buyer) Buy(s Shop) Buyer {
	if b.money > s.price {
		b.money -= s.price
	}

	return b
}

func (s Shop) Sell(b Buyer) Shop {
	if b.money > s.price {
		b.money -= s.price
	}
	return s
}
