package main

import "fmt"

// a function that will show the number of cars that are left to be sold
// function to show the sum of all cars left.
// function to see the number of cars sold.
// a function to show the Sum total of the prices of cars he has sold
// a function to show the list of orders that for the sales he made

// The Car class can have any car attributes you can think of.


type car []string

func (c car) print(){
	for i, singleCar := range c {
		fmt.Println(i, singleCar)
	}
}