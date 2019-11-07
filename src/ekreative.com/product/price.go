package product

type Price struct {
	amount int
}

func MakePrice(amount int) Price {
	return Price{amount: amount}
}

func (p *Price) Amount() int {
	return p.amount
}
