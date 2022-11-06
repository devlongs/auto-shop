package main

import "fmt"

type car product


// a function that will show the number of cars that are left to be sold
func (c car) showNumOfCars(){
	fmt.Println(len(c))
}

// function to show the sum of all cars left.
func (c car) sumOfCarsLeft(){
	var sum int
	for _, singleCar := range c {
		sum += c.price
	}

	fmt.Println(sum)
}

// function to see the number of cars sold.
func (c car) noOfCarsSold(){
	initialNumOfCars := 3
	carsSold := initialNumOfCars - len(c)

	fmt.Println(carsSold)
	
}

// a function to show the Sum total of the prices of cars he has sold
func (c car) sumTotalOfSoldCars(){
	var initialNumOfCars int
	var sum
	for _, singleCar := range c {
		sum += c
	}
}


// a function to show the list of orders that for the sales he made
func (c car) listOfOrders(){
	var sum int
	for _, singleCar := range c {
		sum += c
	}
}