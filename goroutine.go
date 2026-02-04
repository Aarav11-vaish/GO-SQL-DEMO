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
func productWorker(sku int, prodch chan<- product) { // what is this called ? ans : channel send operation
	prodch <- getProduct(sku)
}

// Goroutine function to send price to channel
func priceWorker(sku int, pricech chan<- price) {
	pricech <- getPrice(sku)
}

func simple_goroutineFunc(sku int) (product, price) {
	prodch := make(chan product) // what is this called? ans : channel creation
	pricech := make(chan price)  // what is this called? ans : channel creation

	go productWorker(sku, prodch)
	go priceWorker(sku, pricech)

	prod := <-prodch // what is this called ? ans : channel receive operation
	prc := <-pricech // what is this called ? ans : channel receive operation

	return prod, prc
}

func calling() {
	prod, prc := simple_goroutineFunc(1)
	fmt.Println(prod, prc)
}

// how could i used it in db ? ans : you can use goroutines to perform concurrent database operations, such as inserting data, querying data, or updating records.
// For example, you could create a goroutine for each database operation and
//  use channels to communicate the results back to the main function.
//  This can help improve the performance of your application by allowing
// multiple database operations to run concurrently.
