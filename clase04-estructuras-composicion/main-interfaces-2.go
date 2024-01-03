/*
Ejercico productos
*/

package main

import "fmt"

// Interfaz
type IProduct interface {
	getPrice() float64
}

// Metodo de la interfaz
func getPrice(product IProduct) float64 {
	return product.getPrice()
}

// Estructuras
type Small struct {
	Value        float64
	Tax          float64
	ShippingCost float64
}
type Medium struct {
	Value        float64
	Tax          float64
	ShippingCost float64
}
type Large struct {
	Value        float64
	Tax          float64
	ShippingCost float64
}

// Metodos
func (p Small) getPrice() float64 {
	return p.Value*p.Tax + p.ShippingCost
}
func (p Medium) getPrice() float64 {
	return p.Value*p.Tax + p.ShippingCost
}
func (p Large) getPrice() float64 {
	return p.Value*p.Tax + p.ShippingCost
}

const (
	smallType  = "SMALL"
	mediumType = "MEDIUM"
	largeType  = "LARGE"
)

// Factory
func newProduct(ptype string, baseValue float64) IProduct {
	switch ptype {
	case smallType:
		return Small{Value: baseValue, Tax: 1, ShippingCost: 0}
	case mediumType:
		return Medium{Value: baseValue, Tax: 1.03, ShippingCost: 0}
	case largeType:
		return Large{Value: baseValue, Tax: 1.03, ShippingCost: 2500}
	}
	return nil
}

func main() {

	var p1 = newProduct("SMALL", 1)
	price := p1.getPrice()
	fmt.Printf("Small: %v \n", price)

	var p2 = newProduct("MEDIUM", 1)
	price = p2.getPrice()
	fmt.Printf("Medium: %v \n", price)

	var p3 = newProduct("LARGE", 1)
	price = p3.getPrice()
	fmt.Printf("Large: %v \n", price)

}
