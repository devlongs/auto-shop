// The Product class should have attributes of a product i.e (the product, quantity of the product in stock, price of the product). A car is a product of the store, but there can be other products so the attribute of the car can be promoted to the Product. The Product class should have methods to display a product, and a method to display the status of a product if it is still in stock or not.

package main

import "fmt"

type productProperties struct{
	product string
	quantity int
	price int
	status bool
}

type product productProperties

func (p product) displayProduct(){
	fmt.Printf("%v\n", p)
}

func (p product) displayStatus(){
	fmt.Printf("%v\n", p.status)
}