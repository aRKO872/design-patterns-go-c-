package main

import "fmt"

type ProductPrototype interface {
	Clone() ProductPrototype
	Display() string
}

type Product struct {
	name string
	price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		name: name,
		price: price,
	}
}

func (p *Product) Clone() ProductPrototype {
	return NewProduct(p.name, p.price)
}

func (p *Product) Display() string {
	return "Product Name: " + p.name + ", Price: " + fmt.Sprintf("%.2f", p.price)
}

func main() {
	prod1 := NewProduct("Laptop", 999.99)
	prod2 := NewProduct("Smartphone", 499.99)

	prod1Clone := prod1.Clone()
	prod2Clone := prod2.Clone()

	fmt.Println(prod1.Display())
	fmt.Println(prod2.Display())
	fmt.Println(prod1Clone.Display())
	fmt.Println(prod2Clone.Display())
}