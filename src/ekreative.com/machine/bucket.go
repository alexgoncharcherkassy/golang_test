package machine

import (
	"awesomeProject/src/ekreative.com/product"
	"errors"
)

type Bucket struct {
	elements []product.Product
	size     int
}

func (b Bucket) Elements() []product.Product {
	return b.elements
}

func MakeBucket(size int) Bucket {
	return Bucket{
		elements: []product.Product{},
		size:     size,
	}
}

func (b *Bucket) Push(element product.Product) (*product.Product, error) {
	if len(b.elements) < b.size {
		b.elements = append(b.elements, element)

		return &element, nil
	}

	return nil, errors.New("bucket full")
}

func (b *Bucket) Pop() (*product.Product, error) {
	var element product.Product

	if len(b.elements) == 0 {
		return nil, errors.New("bucket empty")
	}

	element, b.elements = b.elements[0], b.elements[1:]

	return &element, nil
}

func (b *Bucket) IsElemPresentOnFirstPosition(productName int) bool {
	firstProduct := &b.elements[0]

	if firstProduct != nil && firstProduct.Name() == productName {
		return true
	}

	return false
}

func (b *Bucket) IsElemPresentOnPosition(productName int, position int) bool {
	searchingProduct := &b.elements[position]

	if searchingProduct != nil && searchingProduct.Name() == productName {
		return true
	}

	return false
}

func (b *Bucket) GetElemPosition(productName int) []int {
	var positions []int

	for key, element := range b.elements {
		if element.Name() == productName {
			positions = append(positions, key)
		}
	}

	return positions
}

func (b *Bucket) IsElemPresentOnBucket(productName int) bool {
	for _, element := range b.elements {
		if element.Name() == productName {
			return true
		}
	}
	return false
}

func (b *Bucket) GetElemIfPresentOnFirstPosition(productName int) (*product.Product, error) {
	firstProduct := &b.elements[0]

	if b.IsElemPresentOnFirstPosition(productName) {
		return firstProduct, nil
	} else {
		return nil, errors.New("not found")
	}
}

func (b *Bucket) GetElemByNameAdnPosition(productName int, position int) (*product.Product, error) {
	for key, val := range b.elements {
		if key == position && val.Name() == productName {
			return &val, nil
		}
	}

	return nil, errors.New("product not found")
}
