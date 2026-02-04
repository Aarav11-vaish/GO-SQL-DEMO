package main

import "fmt"

type product struct {
	id   int
	name string
}
type price struct {
	SKU    int
	amount int
}

func getProduct(sku int) product {
	return product{
		id:   sku,
		name: "Apple",
	}
}

func getPrice(sku int) price {
	return price{
		SKU:    sku,
		amount: 100,
	}
}

// Goroutine function to send product to channel
func productWorker(sku int, prodch chan<- product) {
	prodch <- getProduct(sku)
}

// Goroutine function to send price to channel
func priceWorker(sku int, pricech chan<- price) {
	pricech <- getPrice(sku)
}

func simple_goroutineFunc(sku int) (product, price) {
	prodch := make(chan product)
	pricech := make(chan price)

	go productWorker(sku, prodch)
	go priceWorker(sku, pricech)

	prod := <-prodch
	prc := <-pricech

	return prod, prc
}

func main() {
	prod, prc := simple_goroutineFunc(1)
	fmt.Println(prod, prc)
}
