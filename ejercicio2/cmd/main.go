package main

import (
	"ejercicio2/internal/product"
	"fmt"
)

func main() {
	p, err := product.Factory("LARGE", 10)

	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("price: ", p.Price())

}
