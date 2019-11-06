package product

type Product struct {
	name int
}

func MakeProduct(name int) Product {
	return Product{name: name}
}

func (p *Product) Name() int {
	return p.name
}
