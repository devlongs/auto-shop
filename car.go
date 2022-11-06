package main

import "fmt"




// a function to show the list of orders that for the sales he made

// The Car class can have any car attributes you can think of.


type car []string

// a function that will show the number of cars that are left to be sold
func (c car) showNumOfCars(){
	fmt.Println(len(c))
}

// function to show the sum of all cars left.
func (c car) sumOfCarsLeft(){
	var sum int
	for _, singleCar := range c {
		sum += c
	}
}

// function to see the number of cars sold.
func (c car) noOfCarsSold(){
	var sum int
	for _, singleCar := range c {
		sum += c
	}
}

// a function to show the Sum total of the prices of cars he has sold
func (c car) sumTotalOfSoldCars(){
	var sum int
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