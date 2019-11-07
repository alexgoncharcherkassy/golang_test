package product

type Product struct {
	name  int
	price Price
}

func MakeProduct(name int, price Price) Product {
	return Product{name: name, price: price}
}

func (p *Product) Name() int {
	return p.name
}

func (p *Product) Price() *Price {
	return &p.price
}

func (p *Product) SetPrice(price Price) *Product {
	p.price = price

	return p
}
